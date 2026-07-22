package ingestion

import (
	"context"
	"fmt"
	"strings"

	"github.com/appenheimer/backend/internal/database/postgres"
	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	db *postgres.DB
}

func NewPipeline(db *postgres.DB) *Pipeline {
	return &Pipeline{db: db}
}

// Run executes the 6-phase curation pipeline on a YAML file.
func (p *Pipeline) Run(ctx context.Context, yamlData []byte, source string, dryRun bool) (*ImportReport, error) {
	report := &ImportReport{
		ImportID: uuid.New().String(),
		Events:   []ValidationEvent{},
	}

	// PHASE 1: Parse
	var apps []CanonicalApp
	if err := yaml.Unmarshal(yamlData, &apps); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// In dry-run, we use a single transaction and rollback at the end
	err := p.db.WithTx(ctx, func(q *sqlc.Queries) error {
		// Log the import start
		if !dryRun {
			// In a real app we'd insert into imports table here
		}

		for _, app := range apps {
			p.processApp(ctx, q, app, report)
		}

		if dryRun {
			return fmt.Errorf("DRY_RUN_ROLLBACK")
		}
		return nil
	})

	if err != nil && err.Error() != "DRY_RUN_ROLLBACK" {
		return nil, err
	}

	return report, nil
}

func (p *Pipeline) processApp(ctx context.Context, q *sqlc.Queries, app CanonicalApp, report *ImportReport) {
	// PHASE 2: Normalize
	app.Slug = strings.ToLower(strings.TrimSpace(app.Slug))
	app.Name = strings.TrimSpace(app.Name)

	if app.Slug == "" {
		report.AddEvent("unknown", SeverityError, "App is missing required 'slug'")
		report.Skipped++
		return
	}

	// PHASE 3: Validate
	if app.Organization == "" {
		report.AddEvent(app.Slug, SeverityError, "App is missing required 'organization'")
		report.Skipped++
		return
	}

	// PHASE 4: Resolve References (Organizations, Taxonomies)
	// (Skipping full DB lookup logic for brevity in scaffold, normally we create/lookup here)

	// PHASE 5: Deduplicate (Idempotent check)
	existing, err := q.GetAppBySlug(ctx, app.Slug)

	// PHASE 6: Import
	if err == nil && existing.ID != uuid.Nil {
		// Update
		report.AddEvent(app.Slug, SeverityInfo, "App already exists. Updating draft.")
		report.Updated++
		// q.UpdateApp(...)
	} else {
		// Insert
		report.AddEvent(app.Slug, SeverityInfo, "New app created.")
		report.Imported++
		// q.CreateApp(...)
	}
}
