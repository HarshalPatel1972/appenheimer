package health

type ComponentStatus string

const (
	Up      ComponentStatus = "up"
	Down    ComponentStatus = "down"
	Unknown ComponentStatus = "unknown"
)

type ReadinessResponse struct {
	Status     string                     `json:"status"` // "ready" or "not_ready"
	Components map[string]ComponentStatus `json:"components"`
	Metadata   BuildMetadata              `json:"metadata"`
}

type BuildMetadata struct {
	Version        string `json:"version"`
	GitSHA         string `json:"git_sha"`
	Branch         string `json:"branch"`
	BuildTime      string `json:"build_time"`
	GoVersion      string `json:"go_version"`
	SchemaVersion  string `json:"schema_version"`
	SearchProvider string `json:"search_provider"`
	Environment    string `json:"environment"`
}

var CurrentMetadata = BuildMetadata{
	Version:        "unknown",
	GitSHA:         "unknown",
	Branch:         "main",
	BuildTime:      "unknown",
	GoVersion:      "unknown",
	SchemaVersion:  "unknown",
	SearchProvider: "postgres",
	Environment:    "development",
}
