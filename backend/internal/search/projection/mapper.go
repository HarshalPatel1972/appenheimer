package projection

type FullAppGraph struct {
	App          interface{}
	Organization interface{}
	Categories   []string
	Platforms    []string
	Pricing      []string
	Tags         []string
	Aliases      []string
}

func BuildDocument(graph FullAppGraph) SearchDocument {
	return SearchDocument{
		ID:            "mock-id",
		Slug:          "mock-slug",
		Name:          "Mock App",
		Description:   "Mock description",
		Organization:  "Mock Org",
		Categories:    graph.Categories,
		Platforms:     graph.Platforms,
		Pricing:       graph.Pricing,
		Tags:          graph.Tags,
		Aliases:       graph.Aliases,
		Status:        "published",
		RankingWeight: 100,
	}
}
