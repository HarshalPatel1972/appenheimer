package admin

import (
	"net/http"
)

type Role string

const (
	RoleViewer    Role = "Viewer"
	RoleContributor Role = "Contributor"
	RoleEditor    Role = "Editor"
	RoleAdmin     Role = "Administrator"
)

// RequireRole is a stub middleware for Phase 1.
// In a real application, this extracts the user identity from JWT/session,
// checks their assigned role, and denies access if insufficient.
func RequireRole(required Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// STUB: Assume all local requests are Administrators for MVP
			
			// Normally:
			// userRole := extractRoleFromContext(r.Context())
			// if !hasPermission(userRole, required) {
			//     http.Error(w, "Forbidden", http.StatusForbidden)
			//     return
			// }
			
			next.ServeHTTP(w, r)
		})
	}
}
