package handlers

import "net/http"

type Handlers struct {
	Handler *http.Handler
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Init() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebsocketHandler)
}
