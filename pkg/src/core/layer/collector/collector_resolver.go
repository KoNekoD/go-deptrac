package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
)

type CollectorResolver struct {
	collectorProvider *CollectorProvider
}

func NewCollectorResolver(collectorProvider *CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(config map[string]interface{}) (*Collectable, error) {
	classLikeType, err := CollectorType.NewCollectorTypeFromString(config["type"].(string))
	if err != nil {
		return nil, err
	}

	if !c.collectorProvider.Has(classLikeType) {
		return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionUnsupportedType(classLikeType, c.collectorProvider.GetKnownCollectors(), nil)
	}

	collector := c.collectorProvider.Get(classLikeType)

	return NewCollectable(collector, config), nil
}
