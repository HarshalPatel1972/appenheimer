package handlers

import (
	"context"
	
	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/appenheimer/backend/internal/search"
	"github.com/appenheimer/backend/internal/search/projection"
)

type AppPublishedHandler struct {
	q        *sqlc.Queries
	provider search.IndexProvider
}

func NewAppPublishedHandler(q *sqlc.Queries, p search.IndexProvider) *AppPublishedHandler {
	return &AppPublishedHandler{q: q, provider: p}
}

func (h *AppPublishedHandler) Handle(ctx context.Context, event sqlc.OutboxEvent) error {
	// Parse aggregate_id and fetch from DB (mocked for compilation)
	graph := projection.FullAppGraph{
		App: nil,
	}
	
	doc := projection.BuildDocument(graph)
	return h.provider.Upsert(ctx, doc)
}

type AppArchivedHandler struct {
	provider search.IndexProvider
}

func NewAppArchivedHandler(p search.IndexProvider) *AppArchivedHandler {
	return &AppArchivedHandler{provider: p}
}

func (h *AppArchivedHandler) Handle(ctx context.Context, event sqlc.OutboxEvent) error {
	return h.provider.Delete(ctx, "mock-id")
}
