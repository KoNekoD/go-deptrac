package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
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
