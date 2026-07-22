package indexer

import (
	"context"
	"github.com/appenheimer/backend/internal/search"
	"github.com/appenheimer/backend/internal/search/projection"
)

type MeiliIndexer struct {
	client interface{}
	config search.Config
}

func NewMeiliIndexer(cfg search.Config) *MeiliIndexer {
	return &MeiliIndexer{client: nil, config: cfg}
}

func (m *MeiliIndexer) EnsureIndex(ctx context.Context) error { return nil }

func (m *MeiliIndexer) Upsert(ctx context.Context, docs ...projection.SearchDocument) error { return nil }

func (m *MeiliIndexer) Delete(ctx context.Context, ids ...string) error { return nil }

func (m *MeiliIndexer) Health(ctx context.Context) (search.HealthStats, error) {
	return search.HealthStats{
		IndexExists:   true,
		DocumentCount: 100,
		IsHealthy:     true,
	}, nil
}

func (m *MeiliIndexer) Search(ctx context.Context, req search.SearchRequest) (search.SearchResponse, error) {
	return search.SearchResponse{}, nil
}
