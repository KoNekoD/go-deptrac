package CollectorResolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/Collectable"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/CollectorProvider"
)

type CollectorResolver struct {
	collectorProvider *CollectorProvider.CollectorProvider
}

func NewCollectorResolver(collectorProvider *CollectorProvider.CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(config map[string]interface{}) (*Collectable.Collectable, error) {
	classLikeType, err := CollectorType.NewCollectorTypeFromString(config["type"].(string))
	if err != nil {
		return nil, err
	}

	if !c.collectorProvider.Has(classLikeType) {
		return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionUnsupportedType(classLikeType, c.collectorProvider.GetKnownCollectors(), nil)
	}

	collector := c.collectorProvider.Get(classLikeType)

	return Collectable.NewCollectable(collector, config), nil
}
