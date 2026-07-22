package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	v1 "github.com/appenheimer/backend/internal/api/v1"
	"github.com/appenheimer/backend/internal/config"
	"github.com/appenheimer/backend/internal/database/postgres"
)

type Server struct {
	httpServer *http.Server
	log        *slog.Logger
}

func New(cfg *config.Config, logger *slog.Logger, db *postgres.DB) *Server {
	mux := http.NewServeMux()

	// Mount V1 routes
	v1.RegisterRoutes(mux, db)

	// Apply Middleware in strict order:
	// Recovery -> Request ID -> Logging -> Compression -> Routes
	handler := Chain(mux,
		Recovery(logger),
		RequestID(),
		Logging(logger),
		Compression(),
	)

	return &Server{
		log: logger,
		httpServer: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.Port),
			Handler:           handler,
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       120 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	s.log.Info("server listening", "addr", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("server shutting down gracefully")
	return s.httpServer.Shutdown(ctx)
}
