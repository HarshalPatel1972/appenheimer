package observability

import (
	"encoding/json"
	"fmt"

	"github.com/appenheimer/backend/internal/observability/health"
)

func GenerateDiagnosticReport() string {
	res := health.ReadinessResponse{
		Status: "ready",
		Components: map[string]health.ComponentStatus{
			"postgres":    health.Up,
			"meilisearch": health.Up,
			"worker":      health.Up,
			"events":      health.Up,
		},
		Metadata: health.CurrentMetadata,
	}

	bytes, _ := json.MarshalIndent(res, "", "  ")
	return fmt.Sprintf("Observability Diagnostics Scaffolded Successfully:\\n%s", string(bytes))
}
