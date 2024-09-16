package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
)

type CollectorResolver struct {
	collectorProvider *CollectorProvider
}

func NewCollectorResolver(collectorProvider *CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(configMap map[string]interface{}) (*Collectable, error) {
	classLikeType, err := config_contract.NewCollectorTypeFromString(configMap["type"].(string))
	if err != nil {
		return nil, err
	}

	if !c.collectorProvider.Has(classLikeType) {
		return nil, layer_contract.NewInvalidCollectorDefinitionExceptionUnsupportedType(classLikeType, c.collectorProvider.GetKnownCollectors(), nil)
	}

	collector := c.collectorProvider.Get(classLikeType)

	return NewCollectable(collector, configMap), nil
}
