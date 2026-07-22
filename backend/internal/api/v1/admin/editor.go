package admin

import (
	"encoding/json"
	"net/http"
	"github.com/appenheimer/backend/internal/database/postgres"
)

type BatchPatchPayload struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	OrganizationID string   `json:"organization_id"`
	HealthStatus   string   `json:"health_status"`
	RankingWeight  int32    `json:"ranking_weight"`
	IconKey        string   `json:"icon_key"`
	Version        int64    `json:"version"` // required for optimistic locking
	
	// Taxonomies
	Categories []string `json:"categories"`
	Platforms  []string `json:"platforms"`
	Pricing    []string `json:"pricing"`
	
	// References
	Aliases []string `json:"aliases"`
	Links   []struct{ Type string `json:"type"`; URL string `json:"url"` } `json:"links"`
	Media   []struct{ Key string `json:"key"`; Type string `json:"type"`; Order int32 `json:"order"` } `json:"media"`
}

func (h *AdminHandler) PatchApp(w http.ResponseWriter, r *http.Request) {
	// 1. Decode payload
	// 2. Validate version exists
	// 3. h.db.WithTx(func(q *sqlc.Queries) error {
	//    a. q.PatchApp (fails if version mismatches)
	//    b. q.DeleteAppAliases, q.CreateAppAlias
	//    c. q.DeleteAppLinks, q.CreateAppLink
	//    d. q.DeleteAppMedia, q.CreateAppMedia
	//    e. Delete/Insert Taxonomies
	// })
	// 4. Return 200 or 409 Conflict
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) CreateShadowDraft(w http.ResponseWriter, r *http.Request) {
	// 1. Get published app
	// 2. Insert new app with source_app_id = published_app.ID, visibility = 'draft', version = 1
	// 3. Copy all relations (aliases, links, media, taxonomies) to new ID
	// 4. Return new draft ID
	w.WriteHeader(http.StatusOK)
}
