package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/appenheimer/backend/internal/search"
)

func HandleGetApp(w http.ResponseWriter, r *http.Request) {
	// Simple path param extraction since we are using standard net/http without a heavy router
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 { // /api/v1/apps/{id}
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}
	id := parts[4]

	details, found := search.GetAppDetails(id)
	if !found {
		http.Error(w, "app not found", http.StatusNotFound)
		return
	}

	resp := AppDetailsResponse{
		AppResponse: AppResponse{
			ID:          details.ID,
			Name:        details.Name,
			Description: details.Description,
			Icon:        details.Icon,
			Categories:  details.Categories,
			Status:      string(details.Status),
		},
		Developer:   details.Developer,
		WebsiteURL:  details.WebsiteURL,
		Pricing:     details.Pricing,
		Platforms:   details.Platforms,
		Screenshots: details.Screenshots,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
