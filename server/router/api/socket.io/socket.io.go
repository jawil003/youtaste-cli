package socket_io

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func RegisterSocketIO(r *gin.RouterGroup) {
	server := socketio.NewServer(nil)

	RegisterPolls(server)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {
			log.Printf("socketio close error: %s\n", err)
		}
	}(server)

	r.GET("/polls/*any", gin.WrapH(server))
	r.POST("/polls/*any", gin.WrapH(server))
}
