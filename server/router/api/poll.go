package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
		/*corsUrlConn := os.Getenv("CORS_URL")
		corsUrl := strings.Split(corsUrlConn, ",")
		res := fmt.Sprintf("http://%s", r.Host)
		return funk.ContainsString(corsUrl, res)*/
		return true
	},
}

func RegisterPolls(r *gin.RouterGroup) {

	pollsGroup := r.Group("/polls")

	hub := observer.NewPollObserverHub()
	go hub.Run()

	pollsGroup.GET("", func(context *gin.Context) {
		polls, err := services.DB().Poll().GetAll()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, polls)
	})

	pollsGroup.GET("/ws", func(context *gin.Context) {
		conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Printf("failed to set ws upgrade: %+v\n", err)
			return
		}

		client := &observer.PollObserverClient{Hub: hub, Conn: conn, Send: make(chan models.Poll, 256)}
		client.Hub.Register <- client

		client.ReadPump()
		client.WritePump()

	})

	pollsGroup.POST("/new", func(context *gin.Context) {
		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		if token == "" {
			context.JSON(400, gin.H{
				"error": "Authorization header is empty",
			})
			return
		}

		jwt := models.Jwt{}

		user, err := services.JWT().Decode(token, &jwt)

		if err != nil || user == nil {
			context.JSON(401, gin.H{
				"message": "Unauthorized",
			})
		}

		var poll models.Poll

		err = context.BindJSON(&poll)

		err = services.DB().Poll().Create(poll, services.User().GetUsername(jwt.Firstname, jwt.Lastname))
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		hub.SendAll(poll)

	})

}
