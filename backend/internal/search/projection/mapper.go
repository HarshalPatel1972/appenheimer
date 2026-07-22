package projection

import (
	"errors"
)

type FullAppGraph struct {
	App          interface{}
	Organization interface{}
	Categories   []string
	Platforms    []string
	Pricing      []string
	Tags         []string
	Aliases      []string
}

func BuildDocument(graph FullAppGraph) (*SearchDocument, error) {
	return nil, errors.New("not implemented: projection mapping")
}
