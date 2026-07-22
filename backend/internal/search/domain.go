package search

type AppStatus string

const (
	Operational AppStatus = "Operational"
	Degraded    AppStatus = "Degraded"
	Maintenance AppStatus = "Maintenance"
	Outage      AppStatus = "Outage"
	Unknown     AppStatus = "Unknown"
)

type App struct {
	ID           string
	Name         string
	Description  string
	Aliases      []string
	Categories   []string
	Pricing      []string
	Platforms    []string
	Icon         string
	SearchWeight int
	Status       AppStatus
}

type AppDetails struct {
	App
	Developer  string
	WebsiteURL string

	Screenshots []string
}

type FilterCriteria struct {
	Platforms  []string
	Pricing    []string
	Categories []string
}

type AvailableFilters struct {
	Platforms  []string
	Pricing    []string
	Categories []string
}

type SearchResult struct {
	Apps             []App
	AvailableFilters AvailableFilters
}
