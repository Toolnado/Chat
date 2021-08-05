package wbserver

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) error {
	s.server = &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
