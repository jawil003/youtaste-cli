package observer

import (
	"bs-to-scrapper/server/models"
	"github.com/gorilla/websocket"
	"log"
)

var instance *PollObserver

func NewPollObserver() *PollObserver {

	if instance == nil {
		instance = &PollObserver{}
	}

	return instance
}

type PollObserver struct {
	connections []*websocket.Conn
	polls       chan models.PollObserverTransport
}

func (po PollObserver) Run() {
	go func() {
		for {
			pl := <-po.polls
			for _, conn := range po.connections {
				err := conn.WriteJSON(pl)
				if err != nil {
					log.Default().Panic(err)
				}
			}
		}
	}()
}

func (po *PollObserver) Connect(connection *websocket.Conn) {
	po.connections = append(po.connections, connection)
}

func (po *PollObserver) Disconnect(connection *websocket.Conn) {
	for i, c := range po.connections {
		if c == connection {
			po.connections = append(po.connections[:i], po.connections[i+1:]...)
			break
		}
	}
}

func (po *PollObserver) SendPoll(poll models.PollObserverTransport) {
	po.polls <- poll
}
