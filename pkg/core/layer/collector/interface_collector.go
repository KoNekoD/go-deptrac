package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
)

type InterfaceCollector struct {
	*AbstractTypeCollector
}

func NewInterfaceCollector() *InterfaceCollector {
	return &InterfaceCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *InterfaceCollector) GetType() ast_map.ClassLikeType {
	return ast_map.TypeInterface
}
