package observer

import "bs-to-scrapper/server/models"

type PollObserverHub struct {
	Register   chan *PollObserverClient
	Unregister chan *PollObserverClient
	Clients    map[*PollObserverClient]bool
}

func NewPollObserverHub() *PollObserverHub {
	return &PollObserverHub{
		Register:   make(chan *PollObserverClient),
		Unregister: make(chan *PollObserverClient),
		Clients:    make(map[*PollObserverClient]bool),
	}
}

func (c *PollObserverHub) Run() {
	for {
		select {
		case client := <-c.Register:
			c.Clients[client] = true
		case client := <-c.Unregister:
			if _, ok := c.Clients[client]; ok {
				close(client.Send)
				delete(c.Clients, client)
			}
		}
	}
}

func (c PollObserverHub) SendAll(poll models.Poll) {
	for client := range c.Clients {
		select {
		case client.Send <- poll:
		default:
			close(client.Send)
			delete(c.Clients, client)
		}
	}
}
