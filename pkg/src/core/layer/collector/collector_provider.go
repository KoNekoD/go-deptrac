package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"golang.org/x/exp/maps"
)

type CollectorProvider struct {
	collectors map[config.CollectorType]layer.CollectorInterface
}

func NewCollectorProvider() *CollectorProvider {
	return &CollectorProvider{}
}

func (p *CollectorProvider) Set(collectors map[config.CollectorType]layer.CollectorInterface) *CollectorProvider {
	p.collectors = collectors
	return p
}

func (p *CollectorProvider) Get(id config.CollectorType) layer.CollectorInterface {
	return p.collectors[id]
}

func (p *CollectorProvider) Has(id config.CollectorType) bool {
	_, ok := p.collectors[id]
	return ok
}

func (p *CollectorProvider) GetKnownCollectors() []config.CollectorType {
	return maps.Keys(p.collectors)
}
