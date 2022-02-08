package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connections = make([]*websocket.Conn, 0)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func RegisterWebsocket(r *gin.RouterGroup) {

	r.GET("/ws", func(context *gin.Context) {
		conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Printf("failed to set ws upgrade: %+v\n", err)
			return
		}

		connections = append(connections, conn)
	})

}
