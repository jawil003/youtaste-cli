package observer

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

func (h *PollObserverHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}
