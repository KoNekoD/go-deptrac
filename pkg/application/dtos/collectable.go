package dtos

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/dependencies_collectors"
)

type Collectable struct {
	Collector  dependencies_collectors.CollectorInterface
	Attributes map[string]interface{}
}

func NewCollectable(collector dependencies_collectors.CollectorInterface, attributes map[string]interface{}) *Collectable {
	return &Collectable{Collector: collector, Attributes: attributes}
}
