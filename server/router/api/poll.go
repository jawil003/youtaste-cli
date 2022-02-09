package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
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

	observer.NewPollObserver().Run()

	pollsGroup.GET("", func(context *gin.Context) {
		conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Printf("failed to set ws upgrade: %+v\n", err)
			return
		}

		//TODO: Replace this with blocking channel
		time.Sleep(10 * time.Second)

		observer.NewPollObserver().Connect(conn)

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

		err = services.DB().Poll().Create(poll)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	})

	pollsGroup.POST("/choose", func(context *gin.Context) {
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

		err = services.DB().Poll().Choose(poll)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	})
}
