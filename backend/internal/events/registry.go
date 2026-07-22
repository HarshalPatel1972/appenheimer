package events

import (
	"context"
	"fmt"
	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
)

type Handler interface {
	Handle(ctx context.Context, event sqlc.OutboxEvent) error
}

type Registry struct {
	handlers map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{handlers: make(map[string]Handler)}
}

func (r *Registry) Register(eventType string, handler Handler) {
	r.handlers[eventType] = handler
}

func (r *Registry) Get(eventType string) (Handler, error) {
	if h, ok := r.handlers[eventType]; ok {
		return h, nil
	}
	return nil, fmt.Errorf("no handler registered for event type: %s", eventType)
}
