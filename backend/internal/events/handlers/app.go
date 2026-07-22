package handlers

import (
	"context"
	"errors"

	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/appenheimer/backend/internal/search"
)

type AppPublishedHandler struct {
	q        *sqlc.Queries
	provider search.IndexProvider
}

func NewAppPublishedHandler(q *sqlc.Queries, p search.IndexProvider) *AppPublishedHandler {
	return &AppPublishedHandler{q: q, provider: p}
}

func (h *AppPublishedHandler) Handle(ctx context.Context, event sqlc.OutboxEvent) error {
	return errors.New("not implemented: event parsing and projection mapping")
}

type AppArchivedHandler struct {
	provider search.IndexProvider
}

func NewAppArchivedHandler(p search.IndexProvider) *AppArchivedHandler {
	return &AppArchivedHandler{provider: p}
}

func (h *AppArchivedHandler) Handle(ctx context.Context, event sqlc.OutboxEvent) error {
	return errors.New("not implemented: event archiving")
}
