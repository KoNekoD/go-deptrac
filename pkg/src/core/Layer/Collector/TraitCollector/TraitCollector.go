package TraitCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/AbstractTypeCollector"
)

type TraitCollector struct {
	*AbstractTypeCollector.AbstractTypeCollector
}

func NewTraitCollector() *TraitCollector {
	return &TraitCollector{
		AbstractTypeCollector: AbstractTypeCollector.NewAbstractTypeCollector(),
	}
}

func (c *TraitCollector) GetType() AstMap.ClassLikeType {
	return AstMap.TypeTrait
}
