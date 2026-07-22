package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/appenheimer/backend/internal/apperrors"
	"github.com/appenheimer/backend/internal/search"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		apperrors.WriteHTTP(w, http.StatusMethodNotAllowed, nil)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apperrors.WriteHTTP(w, http.StatusBadRequest, err)
		return
	}

	query := strings.TrimSpace(req.Query)
	if len(query) > 100 {
		apperrors.WriteHTTP(w, http.StatusBadRequest, nil)
		return
	}
	// Empty query is now allowed!

	var filters *search.FilterCriteria
	if req.Filters != nil {
		filters = &search.FilterCriteria{
			Platforms:  req.Filters.Platforms,
			Pricing:    req.Filters.Pricing,
			Categories: req.Filters.Categories,
		}
	}

	result := search.ExecuteSearch(query, filters)

	var respData []AppResponse
	for _, res := range result.Apps {
		respData = append(respData, AppResponse{
			ID:          res.ID,
			Name:        res.Name,
			Description: res.Description,
			Icon:        res.Icon,
			Categories:  res.Categories,
			Status:      string(res.Status),
		})
	}
	if respData == nil {
		respData = []AppResponse{}
	}

	resp := SearchResponse{
		Data: respData,
		AvailableFilters: AvailableFilters{
			Platforms:  result.AvailableFilters.Platforms,
			Pricing:    result.AvailableFilters.Pricing,
			Categories: result.AvailableFilters.Categories,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
