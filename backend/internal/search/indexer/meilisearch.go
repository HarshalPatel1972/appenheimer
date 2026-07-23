package indexer

import (
	"context"
	"errors"

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

func (m *MeiliIndexer) EnsureIndex(ctx context.Context) error {
	return errors.New("not implemented: meilisearch ensure index")
}

func (m *MeiliIndexer) Upsert(ctx context.Context, docs ...projection.SearchDocument) error {
	return errors.New("not implemented: meilisearch upsert")
}

func (m *MeiliIndexer) Delete(ctx context.Context, ids ...string) error {
	return errors.New("not implemented: meilisearch delete")
}

func (m *MeiliIndexer) Health(ctx context.Context) (search.HealthStats, error) {
	return search.HealthStats{}, errors.New("not implemented: meilisearch health")
}

func (m *MeiliIndexer) Search(ctx context.Context, req search.SearchRequest) (search.SearchResponse, error) {
	return search.SearchResponse{}, errors.New("not implemented: meilisearch search")
}
