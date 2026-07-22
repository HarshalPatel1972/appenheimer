package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleSearch_Contract(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		payload        interface{}
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Valid search request",
			method:         http.MethodPost,
			payload:        SearchRequest{Query: "figma"},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var resp SearchResponse
				if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if resp.Data == nil {
					t.Errorf("expected data array, got nil")
				}
				// Verify schema of first item
				if len(resp.Data) > 0 {
					app := resp.Data[0]
					if app.ID == "" || app.Name == "" || app.Status == "" {
						t.Errorf("schema violation: missing required fields id or name")
					}
				}
			},
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			payload:        nil,
			expectedStatus: http.StatusMethodNotAllowed,
			checkResponse:  func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
		{
			name:           "Empty query",
			method:         http.MethodPost,
			payload:        SearchRequest{Query: "   "},
			expectedStatus: http.StatusOK,
			checkResponse:  func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
		{
			name:           "Query too long",
			method:         http.MethodPost,
			payload:        SearchRequest{Query: string(make([]byte, 101))},
			expectedStatus: http.StatusBadRequest,
			checkResponse:  func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
		{
			name:           "Invalid JSON payload",
			method:         http.MethodPost,
			payload:        "invalid json",
			expectedStatus: http.StatusBadRequest,
			checkResponse:  func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			if str, ok := tt.payload.(string); ok {
				body = []byte(str)
			} else if tt.payload != nil {
				body, _ = json.Marshal(tt.payload)
			}

			req, _ := http.NewRequest(tt.method, "/api/v1/search", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			HandleSearch(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}

			tt.checkResponse(t, rr)
		})
	}
}
