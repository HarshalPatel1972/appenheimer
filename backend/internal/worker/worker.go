package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/appenheimer/backend/internal/database/postgres"
	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/appenheimer/backend/internal/events"
)

type Worker struct {
	db       *postgres.DB
	registry *events.Registry
	id       string
	wg       sync.WaitGroup
	quit     chan struct{}
}

func NewWorker(db *postgres.DB, registry *events.Registry, id string) *Worker {
	return &Worker{
		db:       db,
		registry: registry,
		id:       id,
		quit:     make(chan struct{}),
	}
}

func (w *Worker) Start(ctx context.Context) {
	w.wg.Add(1)
	go w.pollLoop(ctx)
	
	w.wg.Add(1)
	go w.reclaimLoop(ctx)
}

func (w *Worker) Stop() {
	close(w.quit)
	w.wg.Wait()
	log.Println("Worker shutdown gracefully.")
}

func (w *Worker) pollLoop(ctx context.Context) {
	defer w.wg.Done()
	
	for {
		select {
		case <-w.quit:
			return
		case <-ctx.Done():
			return
		default:
			processed := w.processBatch(ctx)
			if processed > 0 {
				time.Sleep(100 * time.Millisecond) // Busy queue
			} else {
				time.Sleep(1 * time.Second) // Idle queue
			}
		}
	}
}

func (w *Worker) processBatch(ctx context.Context) int {
	var events []sqlc.OutboxEvent
	
	// Transaction to claim events
	err := w.db.WithTx(ctx, func(q *sqlc.Queries) error {
		claimed, err := q.ClaimEvents(ctx, sqlc.ClaimEventsParams{
			LockedBy: &w.id,
			Limit:    50,
		})
		if err != nil {
			return err
		}
		events = claimed
		return nil
	})
	
	if err != nil || len(events) == 0 {
		return 0
	}
	
	// Process outside the transaction
	for _, event := range events {
		w.processSingleEvent(ctx, event)
	}
	
	return len(events)
}

func (w *Worker) processSingleEvent(ctx context.Context, event sqlc.OutboxEvent) {
	handler, err := w.registry.Get(event.EventType)
	if err != nil {
		w.handleFailure(ctx, event, err)
		return
	}
	
	err = handler.Handle(ctx, event)
	if err != nil {
		w.handleFailure(ctx, event, err)
		return
	}
	
	w.db.Queries().MarkEventCompleted(ctx, event.ID)
	log.Printf("[worker:%s] processed event:%s type:%s", w.id, event.ID, event.EventType)
}

func (w *Worker) handleFailure(ctx context.Context, event sqlc.OutboxEvent, err error) {
	maxRetries := int32(5)
	if event.Attempts >= maxRetries {
		w.db.Queries().MarkEventDead(ctx, sqlc.MarkEventDeadParams{
			ID:        event.ID,
			LastError: &err.Error(),
		})
		log.Printf("[worker:%s] DEAD event:%s type:%s err:%v", w.id, event.ID, event.EventType, err)
		return
	}

	// Exponential backoff
	backoff := time.Duration(1<<event.Attempts) * 5 * time.Second
	nextAttempt := time.Now().Add(backoff)
	
	w.db.Queries().MarkEventFailed(ctx, sqlc.MarkEventFailedParams{
		ID:            event.ID,
		LastError:     &err.Error(),
		NextAttemptAt: nextAttempt,
	})
	log.Printf("[worker:%s] FAILED event:%s type:%s attempt:%d err:%v", w.id, event.ID, event.EventType, event.Attempts+1, err)
}

func (w *Worker) reclaimLoop(ctx context.Context) {
	defer w.wg.Done()
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-w.quit:
			return
		case <-ticker.C:
			w.db.Queries().ReclaimStaleLocks(ctx, 5)
		}
	}
}
