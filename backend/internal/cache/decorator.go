package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"
	"time"
)

// Mock Types
type SearchRequest struct {
	Query      string
	Platforms  []string
	Categories []string
}
type SearchResponse struct {
	Results []string
}
type SearchProvider interface {
	Search(ctx context.Context, req SearchRequest) (SearchResponse, error)
}

type SearchCacheDecorator struct {
	inner SearchProvider
	cache Cache
	ttl   time.Duration
}

func NewSearchCacheDecorator(inner SearchProvider, cache Cache, ttl time.Duration) *SearchCacheDecorator {
	return &SearchCacheDecorator{inner: inner, cache: cache, ttl: ttl}
}

func (d *SearchCacheDecorator) Search(ctx context.Context, req SearchRequest) (SearchResponse, error) {
	key := generateCacheKey(req)

	cached, err := d.cache.Get(ctx, key)
	if err == nil && cached != nil {
		var res SearchResponse
		if err := json.Unmarshal(cached, &res); err == nil {
			return res, nil
		}
	}

	res, err := d.inner.Search(ctx, req)
	if err != nil {
		return res, err
	}

	bytes, _ := json.Marshal(res)
	_ = d.cache.Set(ctx, key, bytes, d.ttl)

	return res, nil
}

func generateCacheKey(req SearchRequest) string {
	q := strings.ToLower(strings.TrimSpace(req.Query))

	// Deduplicate and Sort Arrays
	plats := dedupeAndSort(req.Platforms)
	cats := dedupeAndSort(req.Categories)

	canonical := struct {
		Q string
		P []string
		C []string
	}{
		Q: q,
		P: plats,
		C: cats,
	}

	bytes, _ := json.Marshal(canonical)
	hash := sha256.Sum256(bytes)
	return "search:query:" + hex.EncodeToString(hash[:])
}

func dedupeAndSort(arr []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range arr {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	sort.Strings(res)
	return res
}
