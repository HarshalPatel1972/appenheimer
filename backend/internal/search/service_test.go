package search

import (
	"testing"
)

func TestExecuteSearch(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []string // IDs expected in order
	}{
		{
			name:     "Exact name match",
			query:    "Figma",
			expected: []string{"app-figma", "app-penpot"}, // penpot matches via alias "open source figma"
		},
		{
			name:     "Case insensitive match",
			query:    "fIgMa",
			expected: []string{"app-figma", "app-penpot"},
		},
		{
			name:     "Alias match",
			query:    "vscode",
			expected: []string{"app-vscode"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := ExecuteSearch(tt.query, nil).Apps
			if len(results) != len(tt.expected) {
				t.Errorf("expected %d results, got %d", len(tt.expected), len(results))
				for _, r := range results {
					t.Logf("Got result: %s", r.ID)
				}
			}
			for i, id := range tt.expected {
				if i < len(results) && results[i].ID != id {
					t.Errorf("at index %d expected %s, got %s", i, id, results[i].ID)
				}
			}
		})
	}
}
