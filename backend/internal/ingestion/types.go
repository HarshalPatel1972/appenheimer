package ingestion

type CanonicalApp struct {
	Name         string   `json:"name" yaml:"name"`
	Slug         string   `json:"slug" yaml:"slug"`
	Organization string   `json:"organization" yaml:"organization"`
	Website      string   `json:"website" yaml:"website"`
	Categories   []string `json:"categories" yaml:"categories"`
	Platforms    []string `json:"platforms" yaml:"platforms"`
	Pricing      []string `json:"pricing" yaml:"pricing"`
	Aliases      []string `json:"aliases" yaml:"aliases"`
	Links        map[string]string `json:"links" yaml:"links"`
	Media        []string `json:"media" yaml:"media"`
}
