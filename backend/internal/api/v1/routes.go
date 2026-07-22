package v1

import (
	"net/http"

	"github.com/appenheimer/backend/internal/api/v1/admin"
	"github.com/appenheimer/backend/internal/database/postgres"
)

// RegisterRoutes attaches all v1 endpoints to the mux
func RegisterRoutes(mux *http.ServeMux, db *postgres.DB) {
	// Health endpoints
	mux.HandleFunc("GET /health/live", handleLive)
	mux.HandleFunc("GET /health/ready", handleReady)

	mux.HandleFunc("POST /api/v1/search", HandleSearch)

	adminHandler := admin.NewAdminHandler(db)
	
	// Helper to wrap handlers
	secure := func(h http.HandlerFunc) http.HandlerFunc {
		return admin.RequireRole(admin.RoleAdmin)(h).ServeHTTP
	}

	mux.HandleFunc("GET /api/v1/admin/dashboard", secure(adminHandler.Dashboard))
	mux.HandleFunc("GET /api/v1/admin/queue", secure(adminHandler.Queue))
	mux.HandleFunc("POST /api/v1/admin/review", secure(adminHandler.Review))
	mux.HandleFunc("POST /api/v1/admin/promote", secure(adminHandler.Promote))
	mux.HandleFunc("POST /api/v1/admin/reject", secure(adminHandler.Reject))
	mux.HandleFunc("POST /api/v1/admin/archive", secure(adminHandler.Archive))
	mux.HandleFunc("PATCH /api/v1/admin/app", secure(adminHandler.PatchApp))
	mux.HandleFunc("POST /api/v1/admin/shadow-draft", secure(adminHandler.CreateShadowDraft))
}
