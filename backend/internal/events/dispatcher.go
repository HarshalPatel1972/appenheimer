package events

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
)

type Dispatcher struct {
	q *sqlc.Queries
}

func NewDispatcher(q *sqlc.Queries) *Dispatcher {
	return &Dispatcher{q: q}
}

func (d *Dispatcher) Publish(ctx context.Context, eventType string, aggregateID uuid.UUID, payload interface{}) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return d.q.CreateOutboxEvent(ctx, sqlc.CreateOutboxEventParams{
		ID:          uuid.New(),
		EventType:   eventType,
		Payload:     payloadBytes,
		Status:      "pending",
	})
}
