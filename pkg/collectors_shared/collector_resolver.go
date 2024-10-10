package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type CollectorResolver struct {
	collectorProvider *CollectorProvider
}

func NewCollectorResolver(collectorProvider *CollectorProvider) *CollectorResolver {
	return &CollectorResolver{collectorProvider: collectorProvider}
}

func (c *CollectorResolver) Resolve(configMap map[string]interface{}) (*violations.Collectable, error) {
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

	return violations.NewCollectable(collector, configMap), nil
}
