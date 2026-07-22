package logger

import (
	"log/slog"
	"os"
)

// Init creates a new structured logger.
func Init() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	// In production, this might be replaced with slog.NewJSONHandler
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
