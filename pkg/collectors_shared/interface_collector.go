package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type InterfaceCollector struct {
	*AbstractTypeCollector
}

func NewInterfaceCollector() *InterfaceCollector {
	return &InterfaceCollector{
		AbstractTypeCollector: NewAbstractTypeCollector(),
	}
}

func (c *InterfaceCollector) GetType() enums.ClassLikeType {
	return enums.TypeInterface
}
