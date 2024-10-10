package violations

import "github.com/KoNekoD/go-deptrac/pkg/collectors_shared"

type Collectable struct {
	Collector  collectors_shared.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector collectors_shared.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
