package dtos

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
)

type Collectable struct {
	Collector  services.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector services.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
