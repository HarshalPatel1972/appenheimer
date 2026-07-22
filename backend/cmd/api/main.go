package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/appenheimer/backend/internal/config"
	"github.com/appenheimer/backend/internal/database/postgres"
	"github.com/appenheimer/backend/internal/events"
	"github.com/appenheimer/backend/internal/search"
	"github.com/appenheimer/backend/internal/search/indexer"
	"github.com/appenheimer/backend/internal/server"
	"github.com/appenheimer/backend/internal/worker"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// 1. Configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2. Database
	db, err := postgres.New(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("connected to postgres database")

	// 3. Search Provider
	var searchProvider search.SearchProvider
	if cfg.SearchProvider == "meilisearch" {
		meiliCfg := search.Config{
			Provider:  cfg.SearchProvider,
			Host:      cfg.SearchHost,
			APIKey:    cfg.SearchAPIKey,
			IndexName: cfg.SearchIndex,
		}
		meili := indexer.NewMeiliIndexer(meiliCfg)
		stats, err := meili.Health(ctx)
		if err != nil || !stats.IsHealthy {
			logger.Error("meilisearch unhealthy", "error", err)
		} else {
			if err := meili.EnsureIndex(ctx); err != nil {
				logger.Error("meilisearch index configuration failed", "error", err)
			}
			logger.Info("connected to meilisearch")
			searchProvider = meili
		}
	}

	if searchProvider == nil {
		// Fallback to postgres search
		logger.Info("using postgres search provider")
		// Postgres search provider is instantiated directly or handled internally
		// (Assuming handlers will fallback to db.Queries().SearchApps)
	}

	// 4. Events & Worker
	registry := events.NewRegistry()
	// Register handlers here if needed
	// registry.Register("app.published", handlers.NewAppPublishedHandler(...))

	wrkr := worker.NewWorker(db, registry, "worker-1")
	workerCtx, workerCancel := context.WithCancel(context.Background())
	defer workerCancel()
	wrkr.Start(workerCtx)
	logger.Info("background worker started")

	// 5. HTTP Server
	srv := server.New(cfg, logger, db)

	go func() {
		if err := srv.Start(); err != nil {
			logger.Error("http server stopped", "error", err)
		}
	}()

	// 6. Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("graceful shutdown initiated")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}

	workerCancel()
	wrkr.Stop()

	logger.Info("graceful shutdown complete")
}
