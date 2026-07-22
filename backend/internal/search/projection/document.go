package projection

type SearchDocument struct {
	ID            string   `json:"id"`
	Slug          string   `json:"slug"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Organization  string   `json:"organization"`
	Categories    []string `json:"categories"`
	Platforms     []string `json:"platforms"`
	Pricing       []string `json:"pricing"`
	Tags          []string `json:"tags"`
	Aliases       []string `json:"aliases"`
	Status        string   `json:"status"`
	RankingWeight int32    `json:"ranking_weight"`
}
