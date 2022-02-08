package socket_io

import (
	socketio "github.com/googollee/go-socket.io"
)

func RegisterPolls(server *socketio.Server) {
	server.OnEvent("poll", "vote", func(s socketio.Conn, msg string) string {
		return "vote"
	})
}
