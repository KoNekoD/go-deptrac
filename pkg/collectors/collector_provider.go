package collectors

import (
	"golang.org/x/exp/maps"
)

type CollectorProvider struct {
	collectors map[CollectorType]CollectorInterface
}

func NewCollectorProvider() *CollectorProvider {
	return &CollectorProvider{}
}

func (p *CollectorProvider) Set(collectors map[CollectorType]CollectorInterface) *CollectorProvider {
	p.collectors = collectors
	return p
}

func (p *CollectorProvider) Get(id CollectorType) CollectorInterface {
	return p.collectors[id]
}

func (p *CollectorProvider) Has(id CollectorType) bool {
	_, ok := p.collectors[id]
	return ok
}

func (p *CollectorProvider) GetKnownCollectors() []CollectorType {
	return maps.Keys(p.collectors)
}
