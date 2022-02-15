package api

import (
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/services"
	"errors"
	"fmt"
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

func RegisterProgress(api *gin.RouterGroup) {

	hub := observer.ProgressObserverHub{}

	go hub.Run()

	progress := api.Group("/progress")
	{
		progress.GET("", func(context *gin.Context) {
			treeService := services.DB().ProgressTree()

			context.JSON(200, gin.H{
				"progress": treeService.Tree.Root.Value,
			})

		})

		progress.GET("/ws", func(context *gin.Context) {
			conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
			if err != nil {
				fmt.Printf("failed to set ws upgrade: %+v\n", err)
				return
			}

			client := &observer.ProgressObserverClient{
				Conn: conn,
				Hub:  &hub,
				Send: make(chan string, 256),
			}

			hub.Register <- client

			go client.ReadPump()
			go client.WritePump()

		})

		progress.PUT("", func(context *gin.Context) {
			treeService := services.DB().ProgressTree()

			if treeService.Tree.Root.Steps == nil || len(treeService.Tree.Root.Steps) == 0 {
				context.JSON(400, gin.H{"error": errors.New("no steps left").Error()})
			}

			tree, err := treeService.Next(treeService.Tree.Root.Steps[0].Value)

			if err != nil {
				return
			}

			context.JSON(200, gin.H{
				"progress": tree.Root.Value,
			})

		})
	}
}
