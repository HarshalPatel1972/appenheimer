package postgres

import (
	"context"

	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
	"github.com/google/uuid"
)

type OrganizationsRepository struct {
	db *DB
}

func NewOrganizationsRepository(db *DB) *OrganizationsRepository {
	return &OrganizationsRepository{db: db}
}

func (r *OrganizationsRepository) Create(ctx context.Context, params sqlc.CreateOrganizationParams) (sqlc.Organization, error) {
	return r.db.Queries().CreateOrganization(ctx, params)
}

func (r *OrganizationsRepository) Get(ctx context.Context, id uuid.UUID) (sqlc.Organization, error) {
	return r.db.Queries().GetOrganization(ctx, id)
}
