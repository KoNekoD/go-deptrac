package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type ClassCollector struct {
	*AbstractTypeCollector
}

func NewClassCollector() *ClassCollector {
	return &ClassCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *ClassCollector) GetType() ast_map.ClassLikeType {
	return ast_map.TypeClass
}
