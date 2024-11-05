package dependencies_collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassLikeCollector struct {
	*AbstractTypeCollector
}

func NewClassLikeCollector() *ClassLikeCollector {
	return &ClassLikeCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassLikeCollector) GetType() enums.ClassLikeType {
	return enums.TypeClasslike
}
