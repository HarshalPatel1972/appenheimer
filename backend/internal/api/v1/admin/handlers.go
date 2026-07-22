package admin

import (
	"net/http"

	"github.com/appenheimer/backend/internal/database/postgres"
)

type AdminHandler struct {
	db *postgres.DB
}

func NewAdminHandler(db *postgres.DB) *AdminHandler {
	return &AdminHandler{db: db}
}

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) Queue(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) Review(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) Promote(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) Reject(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func (h *AdminHandler) Archive(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}
