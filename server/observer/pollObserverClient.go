package observer

import (
	"bs-to-scrapper/server/models"
	"github.com/gorilla/websocket"
)

type PollObserverClient struct {
	Hub  *PollObserverHub
	Conn *websocket.Conn
	Send chan []models.Poll
}
