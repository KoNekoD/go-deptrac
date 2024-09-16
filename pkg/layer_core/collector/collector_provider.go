package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"golang.org/x/exp/maps"
)

type CollectorProvider struct {
	collectors map[config_contract.CollectorType]layer_contract.CollectorInterface
}

func NewCollectorProvider() *CollectorProvider {
	return &CollectorProvider{}
}

func (p *CollectorProvider) Set(collectors map[config_contract.CollectorType]layer_contract.CollectorInterface) *CollectorProvider {
	p.collectors = collectors
	return p
}

func (p *CollectorProvider) Get(id config_contract.CollectorType) layer_contract.CollectorInterface {
	return p.collectors[id]
}

func (p *CollectorProvider) Has(id config_contract.CollectorType) bool {
	_, ok := p.collectors[id]
	return ok
}

func (p *CollectorProvider) GetKnownCollectors() []config_contract.CollectorType {
	return maps.Keys(p.collectors)
}
