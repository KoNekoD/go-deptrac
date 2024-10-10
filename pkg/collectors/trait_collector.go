package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/types"
)

type TraitCollector struct {
	*AbstractTypeCollector
}

func NewTraitCollector() *TraitCollector {
	return &TraitCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *TraitCollector) GetType() types.ClassLikeType {
	return types.TypeTrait
}
