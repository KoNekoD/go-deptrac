package dependencies_collectors

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

type LayerCollector struct {
	resolver layers_resolvers.LayerResolverInterface
	resolved map[string]map[string]*bool
}

func NewLayerCollector(resolver layers_resolvers.LayerResolverInterface) *LayerCollector {
	return &LayerCollector{
		resolver: resolver,
	}
}

func (c *LayerCollector) Satisfy(config map[string]interface{}, reference tokens_references.TokenReferenceInterface) (bool, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("LayerCollector needs the layer_contract configuration, expected 'value' config_contract is missing or invalid.")
		}
	}

	configValueLayer := config["value"].(string)

	hasInResolver, err := c.resolver.Has(configValueLayer)
	if err != nil {
		return false, err
	}
	if !hasInResolver {
		return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration(fmt.Sprintf("Unknown layer_contract \"%s\" specified in collector.", configValueLayer))
	}

	token := reference.GetToken().ToString()

	if utils.MapKeyExists(c.resolved, token) && utils.MapKeyExists(c.resolved[token], configValueLayer) {
		if c.resolved[token][configValueLayer] == nil {
			return false, apperrors.NewInvalidLayerDefinitionExceptionCircularTokenReference(token)
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
