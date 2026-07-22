package auth

import (
	"net/http"

	"github.com/appenheimer/backend/internal/rbac"
)

// Pipeline: Recovery -> RequestID -> Logging -> LoadSession -> LoadUser -> Authorization -> Handler

func LoadSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
	})
}

func LoadUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
	})
}

func RequirePermission(perm rbac.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract user from ctx, get Role, call rbac.HasPermission()
			// If false, http.Error(w, "Forbidden", 403)
			next.ServeHTTP(w, r)
		})
	}
}

func EnforceCSRF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch || r.Method == http.MethodDelete {
			// verify token
		}
		next.ServeHTTP(w, r)
	})
}
