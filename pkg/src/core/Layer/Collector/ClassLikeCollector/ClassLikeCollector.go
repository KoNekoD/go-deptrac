package ClassLikeCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/AbstractTypeCollector"
)

type ClassLikeCollector struct {
	*AbstractTypeCollector.AbstractTypeCollector
}

func NewClassLikeCollector() *ClassLikeCollector {
	return &ClassLikeCollector{
		AbstractTypeCollector: AbstractTypeCollector.NewAbstractTypeCollector(),
	}
}

func (c *ClassLikeCollector) GetType() AstMap.ClassLikeType {
	return AstMap.TypeClasslike
}
