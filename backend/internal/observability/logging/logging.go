package logging

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

// Standardized levels: DEBUG, INFO, WARN, ERROR, FATAL

func NewContextHandler() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	// A real implementation would wrap the handler to extract trace.SpanFromContext(ctx).SpanContext().TraceID()
	return slog.New(handler)
}

func Info(ctx context.Context, msg string, args ...any) {
	spanCtx := trace.SpanFromContext(ctx).SpanContext()
	traceID := spanCtx.TraceID().String()

	// Inject trace_id safely
	combined := append(args, slog.String("trace_id", traceID))
	slog.InfoContext(ctx, msg, combined...)
}
