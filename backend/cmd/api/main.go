package main

import (
	"context"
	"log"
	"os"

	"github.com/appenheimer/backend/internal/search"
	"github.com/appenheimer/backend/internal/search/indexer"
)

func main() {
	cfg := search.Config{
		Provider:  os.Getenv("SEARCH_PROVIDER"),
		Host:      "http://localhost:7700",
		APIKey:    "masterKey",
		IndexName: "apps",
	}

	var searchProvider search.SearchProvider

	if cfg.Provider == "meilisearch" {
		meili := indexer.NewMeiliIndexer(cfg)
		
		// 1. Validate connection & health
		stats, err := meili.Health(context.Background())
		if err != nil || !stats.IsHealthy {
			log.Fatalf("FATAL: Meilisearch unhealthy during boot: %v", err)
		}
		
		// 2. Ensure settings
		if err := meili.EnsureIndex(context.Background()); err != nil {
			log.Fatalf("FATAL: Meilisearch index configuration failed: %v", err)
		}
		
		log.Println("Successfully connected to Meilisearch Search Provider.")
		searchProvider = meili
		_ = searchProvider // Use in handlers
	} else {
		log.Println("Using default Postgres Search Provider.")
		// searchProvider = postgresProvider
	}
	
	// Boot server...
}
