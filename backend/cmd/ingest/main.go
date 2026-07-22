package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/appenheimer/backend/internal/database/postgres"
	"github.com/appenheimer/backend/internal/ingestion"
)

func main() {
	var file string
	var dryRun bool
	var source string

	flag.StringVar(&file, "file", "", "Path to the canonical YAML file")
	flag.BoolVar(&dryRun, "dry-run", false, "Run pipeline without committing to the database")
	flag.StringVar(&source, "source", "manual-cli", "Source identifier for this import run")
	flag.Parse()

	if file == "" {
		log.Fatal("Error: --file argument is required")
	}

	yamlData, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// In a real scenario, this comes from ENV
	dbUrl := "postgres://postgres:postgres@localhost:5432/appenheimer?sslmode=disable"
	ctx := context.Background()

	db, err := postgres.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	pipeline := ingestion.NewPipeline(db)
	
	fmt.Printf("Starting Import Pipeline... [Dry Run: %v]\n", dryRun)
	report, err := pipeline.Run(ctx, yamlData, source, dryRun)
	if err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}

	// Output report as JSON
	reportJSON, _ := json.MarshalIndent(report, "", "  ")
	fmt.Printf("\n--- Import Report ---\n%s\n", string(reportJSON))
}
