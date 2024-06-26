package server

import (
	"log"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(mux *http.ServeMux) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s *APIServer) Start() error {
	log.Println("Starting API server...")

	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
