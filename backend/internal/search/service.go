package search

import (
	"sort"
	"strings"
)

func intersect(appVals, filterVals []string) bool {
	if len(filterVals) == 0 {
		return true // no filter applied
	}
	for _, f := range filterVals {
		fLower := strings.ToLower(f)
		for _, v := range appVals {
			if strings.ToLower(v) == fLower {
				return true
			}
		}
	}
	return false
}

func uniqueStrings(in []string, newVals []string) []string {
	m := make(map[string]bool)
	for _, v := range in {
		m[v] = true
	}
	var out []string
	for _, v := range in {
		out = append(out, v)
	}
	for _, nv := range newVals {
		if !m[nv] {
			m[nv] = true
			out = append(out, nv)
		}
	}
	sort.Strings(out)
	return out
}

func tokenize(s string) []string {
	s = strings.ToLower(strings.TrimSpace(s))
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_' || r == '/'
	})
	var out []string
	for _, p := range parts {
		if len(p) > 0 {
			out = append(out, p)
		}
	}
	return out
}

func scoreApp(app App, tokens []string) int {
	if len(tokens) == 0 {
		return 1
	}
	nameLower := strings.ToLower(app.Name)
	descLower := strings.ToLower(app.Description)

	score := 0
	for _, token := range tokens {
		matched := false
		// Name match scores highest
		if strings.Contains(nameLower, token) {
			score += 3
			matched = true
		}
		// Description match
		if strings.Contains(descLower, token) {
			score += 2
			if !matched {
				matched = true
			}
		}
		// Alias match
		for _, alias := range app.Aliases {
			if strings.Contains(strings.ToLower(alias), token) {
				score += 2
				if !matched {
					matched = true
				}
				break
			}
		}
		// Category match
		for _, cat := range app.Categories {
			if strings.Contains(strings.ToLower(cat), token) {
				score += 1
				if !matched {
					matched = true
				}
				break
			}
		}
		if !matched {
			return -1 // this token matched nothing — exclude app entirely
		}
	}
	return score
}

func ExecuteSearch(query string, filters *FilterCriteria) SearchResult {
	tokens := tokenize(query)

	var matched []App

	avail := AvailableFilters{
		Platforms:  []string{},
		Pricing:    []string{},
		Categories: []string{},
	}

	type scored struct {
		app   App
		score int
	}
	var scoredResults []scored

	for _, app := range Dataset {
		// 1. Token-based match scoring
		s := scoreApp(app, tokens)
		if s < 0 {
			continue // at least one token had zero matches
		}

		// 2. Filter Match (OR within group, AND between groups)
		if filters != nil {
			if !intersect(app.Platforms, filters.Platforms) ||
				!intersect(app.Pricing, filters.Pricing) ||
				!intersect(app.Categories, filters.Categories) {
				continue
			}
		}

		// 3. Accumulate Available Filters from matched set
		avail.Platforms = uniqueStrings(avail.Platforms, app.Platforms)
		avail.Pricing = uniqueStrings(avail.Pricing, app.Pricing)
		avail.Categories = uniqueStrings(avail.Categories, app.Categories)

		scoredResults = append(scoredResults, scored{app: app, score: s})
	}

	// 4. Sort by combined score (token relevance) + SearchWeight
	sort.SliceStable(scoredResults, func(i, j int) bool {
		si := scoredResults[i].score*10 + scoredResults[i].app.SearchWeight
		sj := scoredResults[j].score*10 + scoredResults[j].app.SearchWeight
		if si == sj {
			return scoredResults[i].app.ID < scoredResults[j].app.ID
		}
		return si > sj
	})

	for _, r := range scoredResults {
		matched = append(matched, r.app)
	}

	return SearchResult{
		Apps:             matched,
		AvailableFilters: avail,
	}
}
