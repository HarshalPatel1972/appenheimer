package v1

import (
	"net/http"
)

// RegisterRoutes attaches all v1 endpoints to the mux
func RegisterRoutes(mux *http.ServeMux) {
	// Health endpoints
	mux.HandleFunc("GET /health/live", handleLive)
	mux.HandleFunc("GET /health/ready", handleReady)
	
	mux.HandleFunc("POST /api/v1/search", HandleSearch)
}
