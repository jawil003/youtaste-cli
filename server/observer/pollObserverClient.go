package observer

import (
	"bs-to-scrapper/server/models"
	"github.com/gorilla/websocket"
	"log"
)

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
				if err != nil {
					return
				}
				return
			}

			err := c.Conn.WriteJSON(poll)
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
				log.Printf("error: %v", err)
			}
			break
		}
	}
}
