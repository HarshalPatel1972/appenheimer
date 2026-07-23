package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/appenheimer/backend/internal/api/v1"
	"github.com/appenheimer/backend/internal/database/postgres"
)

func TestCaptureHTTPResponses(t *testing.T) {
	mux := http.NewServeMux()

	ctx := context.Background()
	db, _ := postgres.New(ctx, "postgres://dummy:dummy@localhost:5432/dummy")

	v1.RegisterRoutes(mux, db)

	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}

	capture := func(method, path, body string) {
		var req *http.Request
		if body != "" {
			req, _ = http.NewRequest(method, server.URL+path, strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
		} else {
			req, _ = http.NewRequest(method, server.URL+path, nil)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to request %s: %v", path, err)
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		t.Logf("\n=== %s %s ===\nStatus: %d\nBody: %s", method, path, resp.StatusCode, string(respBody))
	}

	capture("POST", "/api/v1/search", `{"query":"figma"}`)
	capture("GET", "/api/v1/admin/dashboard", "")
	capture("GET", "/metrics", "")
}
