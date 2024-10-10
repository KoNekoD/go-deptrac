package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/types"
)

type ClassCollector struct {
	*AbstractTypeCollector
}

func NewClassCollector() *ClassCollector {
	return &ClassCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassCollector) GetType() types.ClassLikeType {
	return types.TypeClass
}
