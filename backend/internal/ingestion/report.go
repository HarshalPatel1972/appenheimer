package ingestion

import "github.com/google/uuid"

type Severity string

const (
	SeverityInfo    Severity = "INFO"
	SeverityWarning Severity = "WARNING"
	SeverityError   Severity = "ERROR"
)

type ValidationEvent struct {
	AppSlug  string   `json:"app_slug"`
	Severity Severity `json:"severity"`
	Message  string   `json:"message"`
}

type ImportReport struct {
	ImportID string            `json:"import_id"`
	Imported int               `json:"imported"`
	Updated  int               `json:"updated"`
	Skipped  int               `json:"skipped"`
	Events   []ValidationEvent `json:"events"`
}

func (r *ImportReport) AddEvent(slug string, sev Severity, msg string) {
	r.Events = append(r.Events, ValidationEvent{
		AppSlug:  slug,
		Severity: sev,
		Message:  msg,
	})
}
