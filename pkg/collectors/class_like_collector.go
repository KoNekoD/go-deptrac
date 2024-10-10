package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/types"
)

type ClassLikeCollector struct {
	*AbstractTypeCollector
}

func NewClassLikeCollector() *ClassLikeCollector {
	return &ClassLikeCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassLikeCollector) GetType() types.ClassLikeType {
	return types.TypeClasslike
}
