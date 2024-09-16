package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type TraitCollector struct {
	*AbstractTypeCollector
}

func NewTraitCollector() *TraitCollector {
	return &TraitCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *TraitCollector) GetType() ast_map.ClassLikeType {
	return ast_map.TypeTrait
}
