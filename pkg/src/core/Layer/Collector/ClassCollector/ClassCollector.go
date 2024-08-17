package ClassCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/AbstractTypeCollector"
)

type ClassCollector struct {
	*AbstractTypeCollector.AbstractTypeCollector
}

func NewClassCollector() *ClassCollector {
	return &ClassCollector{
		AbstractTypeCollector: AbstractTypeCollector.NewAbstractTypeCollector(),
	}
}

func (c *ClassCollector) GetType() AstMap.ClassLikeType {
	return AstMap.TypeClass
}
