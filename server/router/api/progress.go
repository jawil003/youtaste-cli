package api

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterProgress(api *gin.RouterGroup, hub *observer.ProgressObserverHub) {

	go hub.Run()

	progress := api.Group("/progress")
	{
		progress.GET("", func(context *gin.Context) {
			treeService := services.DB().ProgressTree()

			res := gin.H{
				"progress": treeService.Tree.Root.Value,
			}

			logger.Logger().Info.Println(res)
			context.JSON(http.StatusOK, res)

		})

		progress.GET("/ws", func(context *gin.Context) {
			conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
			if err != nil {
				logger.Logger().Error.Println(err)
				return
			}

			client := &observer.ProgressObserverClient{
				Conn: conn,
				Hub:  hub,
				Send: make(chan string, 256),
			}

			hub.Register <- client

			go client.ReadPump()
			go client.WritePump()

			logger.Logger().Info.Printf("New client %s connected", conn.RemoteAddr().String())

		})
	}
}
