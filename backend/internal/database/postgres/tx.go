package postgres

import (
	"context"
	"fmt"

	"github.com/appenheimer/backend/internal/database/postgres/sqlc"
)

func (db *DB) WithTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	q := sqlc.New(tx)
	if err := fn(q); err != nil {
		return err // rollback triggered by defer
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// Queries provides raw query access without a transaction wrapper for single operations
func (db *DB) Queries() *sqlc.Queries {
	return sqlc.New(db.pool)
}
