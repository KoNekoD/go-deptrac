package collector

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CollectorInterface"

type Collectable struct {
	Collector  CollectorInterface.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector CollectorInterface.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
