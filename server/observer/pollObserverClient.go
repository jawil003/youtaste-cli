package observer

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"github.com/gorilla/websocket"
	"log"
)

var infoLogger = logger.Logger().Info
var errorLogger = logger.Logger().Error
var warnLogger = logger.Logger().Warn

type PollObserverClient struct {
	Hub  *PollObserverHub
	Conn *websocket.Conn
	Send chan models.Poll
}

func (c *PollObserverClient) WritePump() {
	defer func() {
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()
	for {
		select {
		case poll, ok := <-c.Send:

			if !ok {
				// The hub closed the channel.
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				warnLogger.Printf("PollObserverClient closed channel %v", c.Conn.RemoteAddr())
				if err != nil {
					return
				}
				return
			}

			infoLogger.Println("PollObserverClient send poll to client %s with args %s", c.Conn.RemoteAddr(), logger.ConvertToString(poll))
			err := c.Conn.WriteJSON(poll)
			log.Printf("PollObserverClient broadcasted: %v", poll)
			if err != nil {
				return
			}

		}
	}
}

func (c *PollObserverClient) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		err := c.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				errorLogger.Println(err)
			}
			break
		}
	}
}
