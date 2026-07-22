package postgres

import (
	"context"

	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AppsRepository struct {
	db *DB
}

func NewAppsRepository(db *DB) *AppsRepository {
	return &AppsRepository{db: db}
}

func (r *AppsRepository) Create(ctx context.Context, params sqlc.CreateAppParams) (sqlc.App, error) {
	return r.db.Queries().CreateApp(ctx, params)
}

func (r *AppsRepository) Get(ctx context.Context, id pgtype.UUID) (sqlc.App, error) {
	return r.db.Queries().GetApp(ctx, id)
}
