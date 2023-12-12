package server

import (
	"context"
	"net/http"

	"gamelib/internal/storage/db"
	"gamelib/pkg/config"
)

type Server struct {
	httpServer *http.Server
	Storage    *db.Storage
}

func NewServer(cfg *config.Config) (*Server, error) {
	storage, err := db.NewStorage(cfg.Storage)
	if err != nil {
		return nil, err
	}

	return &Server{
		Storage: &storage,
	}, nil
}

func (s *Server) Start(Port string) error {
	billingRouter := s.configureRoutes()

	s.httpServer = &http.Server{
		Addr:    ":" + Port,
		Handler: billingRouter,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
