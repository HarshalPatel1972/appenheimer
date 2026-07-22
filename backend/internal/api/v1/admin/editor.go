package admin

import (
	"net/http"
)

type BatchPatchPayload struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organization_id"`
	HealthStatus   string `json:"health_status"`
	RankingWeight  int32  `json:"ranking_weight"`
	IconKey        string `json:"icon_key"`
	Version        int64  `json:"version"` // required for optimistic locking

	// Taxonomies
	Categories []string `json:"categories"`
	Platforms  []string `json:"platforms"`
	Pricing    []string `json:"pricing"`

	// References
	Aliases []string `json:"aliases"`
	Links   []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"links"`
	Media []struct {
		Key   string `json:"key"`
		Type  string `json:"type"`
		Order int32  `json:"order"`
	} `json:"media"`
}

func (h *AdminHandler) PatchApp(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) CreateShadowDraft(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}
