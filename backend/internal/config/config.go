package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           int
	Env            string
	DatabaseURL    string
	SearchProvider string
	SearchHost     string
	SearchAPIKey   string
	SearchIndex    string
}

func Load() (*Config, error) {
	// Attempt to load .env files from various relative paths (ignore errors if not found, since production sets them directly)
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../../.env")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT environment variable: %w", err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	searchProvider := os.Getenv("SEARCH_PROVIDER")
	if searchProvider == "" {
		searchProvider = "postgres" // Use Postgres fallback by default if not specified
	}

	searchHost := os.Getenv("SEARCH_HOST")
	if searchHost == "" {
		searchHost = "http://localhost:7700"
	}

	searchAPIKey := os.Getenv("SEARCH_API_KEY")
	if searchAPIKey == "" {
		searchAPIKey = "masterKey"
	}

	searchIndex := os.Getenv("SEARCH_INDEX")
	if searchIndex == "" {
		searchIndex = "apps"
	}

	return &Config{
		Port:           port,
		Env:            env,
		DatabaseURL:    dbURL,
		SearchProvider: searchProvider,
		SearchHost:     searchHost,
		SearchAPIKey:   searchAPIKey,
		SearchIndex:    searchIndex,
	}, nil
}
