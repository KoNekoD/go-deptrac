package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"golang.org/x/exp/maps"
)

type CollectorProvider struct {
	collectors map[enums.CollectorType]CollectorInterface
}

func NewCollectorProvider() *CollectorProvider {
	return &CollectorProvider{}
}

func (p *CollectorProvider) Set(collectors map[enums.CollectorType]CollectorInterface) *CollectorProvider {
	p.collectors = collectors
	return p
}

func (p *CollectorProvider) Get(id enums.CollectorType) CollectorInterface {
	return p.collectors[id]
}

func (p *CollectorProvider) Has(id enums.CollectorType) bool {
	_, ok := p.collectors[id]
	return ok
}

func (p *CollectorProvider) GetKnownCollectors() []enums.CollectorType {
	return maps.Keys(p.collectors)
}
