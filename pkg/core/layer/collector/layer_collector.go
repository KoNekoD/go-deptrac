package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/core/layer/layer_resolver_interface"
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

func (c *LayerCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return false, layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("LayerCollector needs the layer configuration, expected 'value' config is missing or invalid.")
		}
	}

	configValueLayer := config["value"].(string)

	hasInResolver, err := c.resolver.Has(configValueLayer)
	if err != nil {
		return false, err
	}
	if !hasInResolver {
		return false, layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("Unknown layer \"%s\" specified in collector.", configValueLayer))
	}

	token := reference.GetToken().ToString()

	if util.MapKeyExists(c.resolved, token) && util.MapKeyExists(c.resolved[token], configValueLayer) {
		if c.resolved[token][configValueLayer] == nil {
			return false, layer.NewInvalidLayerDefinitionExceptionCircularTokenReference(token)
		}

		return *c.resolved[token][configValueLayer], nil
	}

	// Set resolved for current token to null in case resolver comes back to it (circular reference)
	c.resolved[token][configValueLayer] = nil

	resolvedValue, err := c.resolver.IsReferenceInLayer(configValueLayer, reference)

	if err != nil {
		return false, err
	}

	c.resolved[token][configValueLayer] = &resolvedValue
	return resolvedValue, nil
}
