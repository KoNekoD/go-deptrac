package InterfaceCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/AbstractTypeCollector"
)

type InterfaceCollector struct {
	*AbstractTypeCollector.AbstractTypeCollector
}

func NewInterfaceCollector() *InterfaceCollector {
	return &InterfaceCollector{
		AbstractTypeCollector: AbstractTypeCollector.NewAbstractTypeCollector(),
	}
}

func (c *InterfaceCollector) GetType() AstMap.ClassLikeType {
	return AstMap.TypeInterface
}
