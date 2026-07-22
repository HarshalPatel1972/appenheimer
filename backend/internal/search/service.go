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

func ExecuteSearch(query string, filters *FilterCriteria) SearchResult {
	query = strings.ToLower(strings.TrimSpace(query))

	var matched []App

	avail := AvailableFilters{
		Platforms:  []string{},
		Pricing:    []string{},
		Categories: []string{},
	}

	for _, app := range Dataset {
		// 1. Text Search Match
		isMatch := query == "" // empty query matches everything
		if !isMatch {
			if strings.Contains(strings.ToLower(app.Name), query) || strings.Contains(strings.ToLower(app.Description), query) {
				isMatch = true
			} else {
				for _, alias := range app.Aliases {
					if strings.Contains(strings.ToLower(alias), query) {
						isMatch = true
						break
					}
				}
				if !isMatch {
					for _, cat := range app.Categories {
						if strings.Contains(strings.ToLower(cat), query) {
							isMatch = true
							break
						}
					}
				}
			}
		}

		if !isMatch {
			continue
		}

		// 2. Filter Match (OR within, AND between)
		if filters != nil {
			if !intersect(app.Platforms, filters.Platforms) ||
				!intersect(app.Pricing, filters.Pricing) ||
				!intersect(app.Categories, filters.Categories) {
				continue
			}
		}

		// 3. Accumulate Available Filters
		avail.Platforms = uniqueStrings(avail.Platforms, app.Platforms)
		avail.Pricing = uniqueStrings(avail.Pricing, app.Pricing)
		avail.Categories = uniqueStrings(avail.Categories, app.Categories)

		matched = append(matched, app)
	}

	// 4. Sort
	sort.SliceStable(matched, func(i, j int) bool {
		if matched[i].SearchWeight == matched[j].SearchWeight {
			return matched[i].ID < matched[j].ID
		}
		return matched[i].SearchWeight > matched[j].SearchWeight
	})

	return SearchResult{
		Apps:             matched,
		AvailableFilters: avail,
	}
}
