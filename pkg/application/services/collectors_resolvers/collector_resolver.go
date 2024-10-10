package collectors_resolvers

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type CollectorResolver struct {
	collectorProvider *services.CollectorProvider
}

func NewCollectorResolver(collectorProvider *services.CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(configMap map[string]interface{}) (*dtos.Collectable, error) {
	classLikeType, err := enums.NewCollectorTypeFromString(configMap["type"].(string))
	if err != nil {
		return nil, err
	}

	if !c.collectorProvider.Has(classLikeType) {
		list := make([]string, 0)

		for _, v := range c.collectorProvider.GetKnownCollectors() {
			list = append(list, string(v))
		}

		return nil, apperrors.NewInvalidCollectorDefinitionUnsupportedType(string(classLikeType), list, nil)
	}

	collector := c.collectorProvider.Get(classLikeType)

	return dtos.NewCollectable(collector, configMap), nil
}
