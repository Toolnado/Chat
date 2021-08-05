package main

import (
	"log"

	"github.com/Toolnado/Chat/internal/handlers"
	"github.com/Toolnado/Chat/internal/wbserver"
)

const port = "8080"

func main() {
	h := handlers.NewHandlers()
	s := wbserver.NewServer()

	h.Init()

	if err := s.Run(port); err != nil {
		log.Fatalf("Server error: %s", err.Error())
		return
	}
}
