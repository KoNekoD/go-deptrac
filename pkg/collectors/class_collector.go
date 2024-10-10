package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassCollector struct {
	*AbstractTypeCollector
}

func NewClassCollector() *ClassCollector {
	return &ClassCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassCollector) GetType() enums.ClassLikeType {
	return enums.TypeClass
}
