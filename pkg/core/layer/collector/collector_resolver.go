package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
)

type CollectorResolver struct {
	collectorProvider *CollectorProvider
}

func NewCollectorResolver(collectorProvider *CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(configMap map[string]interface{}) (*Collectable, error) {
	classLikeType, err := config.NewCollectorTypeFromString(configMap["type"].(string))
	if err != nil {
		return nil, err
	}

	if !c.collectorProvider.Has(classLikeType) {
		return nil, layer.NewInvalidCollectorDefinitionExceptionUnsupportedType(classLikeType, c.collectorProvider.GetKnownCollectors(), nil)
	}

	collector := c.collectorProvider.Get(classLikeType)

	return NewCollectable(collector, configMap), nil
}
