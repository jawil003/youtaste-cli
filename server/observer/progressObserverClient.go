package observer

import (
	"github.com/gorilla/websocket"
	"log"
)

type ProgressObserverClient struct {
	Hub  *ProgressObserverHub
	Conn *websocket.Conn
	Send chan string
}

func (c *ProgressObserverClient) WritePump() {
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
				log.Printf("ProgressObserverClient closed channel %v", c.Conn.RemoteAddr())
				if err != nil {
					return
				}
				return
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, []byte(poll))
			log.Printf("ProgressObserverClient broadcasted: %v", poll)
			if err != nil {
				return
			}

		}
	}
}

func (c *ProgressObserverClient) ReadPump() {
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
