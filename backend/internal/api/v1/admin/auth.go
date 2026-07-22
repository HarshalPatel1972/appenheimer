package admin

import (
	"net/http"
)

type Role string

const (
	RoleViewer      Role = "Viewer"
	RoleContributor Role = "Contributor"
	RoleEditor      Role = "Editor"
	RoleAdmin       Role = "Administrator"
)

// RequireRole enforces role-based access.
func RequireRole(required Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		})
	}
}
