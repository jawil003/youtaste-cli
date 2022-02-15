package observer

type ProgressObserverHub struct {
	Register   chan *ProgressObserverClient
	Unregister chan *ProgressObserverClient
	Clients    map[*ProgressObserverClient]bool
}

func NewProgressObserverHub() *ProgressObserverHub {
	return &ProgressObserverHub{
		Register:   make(chan *ProgressObserverClient),
		Unregister: make(chan *ProgressObserverClient),
		Clients:    make(map[*ProgressObserverClient]bool),
	}
}

func (c *ProgressObserverHub) Run() {
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

func (c ProgressObserverHub) SendAll(message string) {
	for client := range c.Clients {
		select {
		case client.Send <- message:
		default:
			return
		}
	}
}
