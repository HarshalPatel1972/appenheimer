package cache

import (
	"testing"
)

func TestGenerateCacheKey(t *testing.T) {
	req1 := SearchRequest{
		Query:      "  React Native  ",
		Platforms:  []string{"ios", "android", "ios"},
		Categories: []string{"UI"},
	}

	req2 := SearchRequest{
		Query:      "react native",
		Platforms:  []string{"android", "ios"},
		Categories: []string{"UI"},
	}

	key1 := generateCacheKey(req1)
	key2 := generateCacheKey(req2)

	if key1 != key2 {
		t.Errorf("Expected matching keys for equivalent requests, got:\\n%s\\n%s", key1, key2)
	}
}
