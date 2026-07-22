package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("appenheimer-backend")

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return tracer.Start(ctx, name)
}

type TraceContext struct {
	Traceparent string `json:"traceparent"`
	Tracestate  string `json:"tracestate"`
}

// Extract extracts TraceContext from an Outbox Event payload
func ExtractFromEvent(metadata map[string]interface{}) TraceContext {
	// mock extraction
	return TraceContext{}
}
