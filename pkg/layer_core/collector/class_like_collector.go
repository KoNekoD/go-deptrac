package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type ClassLikeCollector struct {
	*AbstractTypeCollector
}

func NewClassLikeCollector() *ClassLikeCollector {
	return &ClassLikeCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassLikeCollector) GetType() ast_map.ClassLikeType {
	return ast_map.TypeClasslike
}
