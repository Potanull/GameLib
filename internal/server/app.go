package server

import (
	"context"
	"github.com/forbiddencoding/howlongtobeat"
	"net/http"

	"gamelib/internal/storage/db"
	"gamelib/pkg/config"
)

type Server struct {
	httpServer *http.Server
	Storage    *db.Storage
	HLTB       *howlongtobeat.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	storage, err := db.NewStorage(cfg.Storage)
	if err != nil {
		return nil, err
	}

	hltb, err := howlongtobeat.New()
	if err != nil {
		return nil, err
	}

	return &Server{
		Storage: &storage,
		HLTB:    hltb,
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
