package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type TraitCollector struct {
	*AbstractTypeCollector
}

func NewTraitCollector() *TraitCollector {
	return &TraitCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *TraitCollector) GetType() enums.ClassLikeType {
	return enums.TypeTrait
}
