package violations

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type Collectable struct {
	Collector  collectors.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector collectors.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
