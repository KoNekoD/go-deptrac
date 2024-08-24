package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
)

type Collectable struct {
	Collector  layer.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector layer.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
