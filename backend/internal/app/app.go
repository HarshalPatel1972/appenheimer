package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/appenheimer/backend/internal/config"
	"github.com/appenheimer/backend/internal/logger"
	"github.com/appenheimer/backend/internal/server"
)

type App struct {
	server *server.Server
	log    *slog.Logger
}

func New() (*App, error) {
	log := logger.Init()

	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// In the future: initialize database and cache connections here.

	srv := server.New(cfg, log)

	return &App{
		server: srv,
		log:    log,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	errChan := make(chan error, 1)

	go func() {
		if err := a.server.Start(); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		a.log.Info("shutdown signal received")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		// In the future: cleanly close database and cache connections here.
		return a.server.Shutdown(shutdownCtx)
	case err := <-errChan:
		return err
	}
}
