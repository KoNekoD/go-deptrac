package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
)

type Collectable struct {
	Collector  layer_contract.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector layer_contract.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
