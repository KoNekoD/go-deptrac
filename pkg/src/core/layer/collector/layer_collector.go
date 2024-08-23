package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidLayerDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type LayerCollector struct {
	resolver layer_resolver_interface.LayerResolverInterface
	resolved map[string]map[string]*bool
}

func NewLayerCollector(resolver layer_resolver_interface.LayerResolverInterface) *LayerCollector {
	return &LayerCollector{
		resolver: resolver,
	}
}

func (c *LayerCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return false, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("LayerCollector needs the layer configuration, expected 'value' config is missing or invalid.")
		}
	}

	layer := config["value"].(string)

	hasInResolver, err := c.resolver.Has(layer)
	if err != nil {
		return false, err
	}
	if !hasInResolver {
		return false, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("Unknown layer \"%s\" specified in collector.", layer))
	}

	token := reference.GetToken().ToString()

	if util.MapKeyExists(c.resolved, token) && util.MapKeyExists(c.resolved[token], layer) {
		if c.resolved[token][layer] == nil {
			return false, InvalidLayerDefinitionException.NewInvalidLayerDefinitionExceptionCircularTokenReference(token)
		}

		return *c.resolved[token][layer], nil
	}

	// Set resolved for current token to null in case resolver comes back to it (circular reference)
	c.resolved[token][layer] = nil

	resolvedValue, err := c.resolver.IsReferenceInLayer(layer, reference)

	if err != nil {
		return false, err
	}

	c.resolved[token][layer] = &resolvedValue
	return resolvedValue, nil
}
