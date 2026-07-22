package search

import (
	"context"
	"github.com/appenheimer/backend/internal/search/projection"
)

type Config struct {
	Provider  string // "postgres" or "meilisearch"
	Host      string
	APIKey    string
	IndexName string
	BatchSize int
}

type HealthStats struct {
	IndexExists   bool
	DocumentCount int64
	PendingTasks  int64
	IsHealthy     bool
}

type SearchRequest struct {
	Query      string
	Platforms  []string
	Categories []string
	Pricing    []string
	Limit      int
	Offset     int
}

type SearchResponse struct {
	Results []projection.SearchDocument
	Total   int64
}

// SearchProvider is injected into public API handlers
type SearchProvider interface {
	Search(ctx context.Context, req SearchRequest) (SearchResponse, error)
}

// IndexProvider is injected into background Outbox workers
type IndexProvider interface {
	Upsert(ctx context.Context, docs ...projection.SearchDocument) error
	Delete(ctx context.Context, ids ...string) error
	Health(ctx context.Context) (HealthStats, error)
	EnsureIndex(ctx context.Context) error
}
