package v1

type SearchFilters struct {
	Platforms  []string `json:"platforms,omitempty"`
	Pricing    []string `json:"pricing,omitempty"`
	Categories []string `json:"categories,omitempty"`
}

type SearchRequest struct {
	Query   string         `json:"query"`
	Filters *SearchFilters `json:"filters,omitempty"`
}

type AvailableFilters struct {
	Platforms  []string `json:"platforms"`
	Pricing    []string `json:"pricing"`
	Categories []string `json:"categories"`
}

type SearchResponse struct {
	Data             []AppResponse    `json:"data"`
	AvailableFilters AvailableFilters `json:"availableFilters"`
}

type AppResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Categories  []string `json:"categories"`
	Status      string   `json:"status"`
}

type AppDetailsResponse struct {
	AppResponse
	Developer   string   `json:"developer"`
	WebsiteURL  string   `json:"websiteUrl"`
	Pricing     []string `json:"pricing"`
	Platforms   []string `json:"platforms"`
	Screenshots []string `json:"screenshots"`
}
