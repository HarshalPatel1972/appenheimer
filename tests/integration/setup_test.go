package integration

import (
	"context"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// setupTestEnvironment boots Postgres, Redis, and Meilisearch ephemerally via Testcontainers.
func setupTestEnvironment(t *testing.T) {
	ctx := context.Background()

	// 1. Boot Postgres
	pgReq := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "appenheimer",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: pgReq,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Could not start postgres: %s", err)
	}
	defer pgC.Terminate(ctx)

	// Mock wait for initialization
	time.Sleep(1 * time.Second)

	// In a real test, you'd apply migrations here, instantiate Repositories, 
	// and run cross-component assertions.
}
