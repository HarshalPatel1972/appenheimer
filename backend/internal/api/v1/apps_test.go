package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetApp_Contract(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Valid app",
			path:           "/api/v1/apps/app-figma",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var resp AppDetailsResponse
				if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if resp.ID != "app-figma" || resp.Name == "" || resp.Developer == "" {
					t.Errorf("schema violation: missing required fields")
				}
			},
		},
		{
			name:           "Not found",
			path:           "/api/v1/apps/unknown-app",
			expectedStatus: http.StatusNotFound,
			checkResponse:  func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()

			HandleGetApp(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}
			tt.checkResponse(t, rr)
		})
	}
}
