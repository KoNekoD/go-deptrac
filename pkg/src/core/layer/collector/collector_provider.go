package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CollectorInterface"
	"golang.org/x/exp/maps"
)

type CollectorProvider struct {
	collectors map[CollectorType.CollectorType]CollectorInterface.CollectorInterface
}

func NewCollectorProvider() *CollectorProvider {
	return &CollectorProvider{}
}

func (p *CollectorProvider) Set(collectors map[CollectorType.CollectorType]CollectorInterface.CollectorInterface) *CollectorProvider {
	p.collectors = collectors
	return p
}

func (p *CollectorProvider) Get(id CollectorType.CollectorType) CollectorInterface.CollectorInterface {
	return p.collectors[id]
}

func (p *CollectorProvider) Has(id CollectorType.CollectorType) bool {
	_, ok := p.collectors[id]
	return ok
}

func (p *CollectorProvider) GetKnownCollectors() []CollectorType.CollectorType {
	return maps.Keys(p.collectors)
}
