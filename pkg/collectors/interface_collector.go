package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/types"
)

type InterfaceCollector struct {
	*AbstractTypeCollector
}

func NewInterfaceCollector() *InterfaceCollector {
	return &InterfaceCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *InterfaceCollector) GetType() types.ClassLikeType {
	return types.TypeInterface
}
