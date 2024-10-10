package analysers

import (
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	layers2 "github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/rules"
)

type RulesetUsageAnalyser struct {
	layerProvider      *services.LayerProvider
	layerResolver      layers2.LayerResolverInterface
	astMapExtractor    *ast_map.AstMapExtractor
	dependencyResolver *pkg.DependencyResolver
	tokenResolver      *services.TokenResolver
	layers             []*rules.Layer
}

func NewRulesetUsageAnalyser(
	layerProvider *services.LayerProvider,
	layerResolver layers2.LayerResolverInterface,
	astMapExtractor *ast_map.AstMapExtractor,
	dependencyResolver *pkg.DependencyResolver,
	tokenResolver *services.TokenResolver,
	layers []*rules.Layer,
) *RulesetUsageAnalyser {
	return &RulesetUsageAnalyser{
		layerProvider:      layerProvider,
		layerResolver:      layerResolver,
		astMapExtractor:    astMapExtractor,
		dependencyResolver: dependencyResolver,
		tokenResolver:      tokenResolver,
		layers:             layers,
	}
}

func (a *RulesetUsageAnalyser) Analyse() (map[string]map[string]int, error) {
	rulesets, err := a.rulesetResolution()
	if err != nil {
		return nil, err
	}

	return a.findRulesetUsages(rulesets)
}

func (a *RulesetUsageAnalyser) rulesetResolution() (map[string]map[string]int, error) {
	rulesets := make(map[string]map[string]int)
	for _, layerDef := range a.layers {
		allowedLayers, err := a.layerProvider.GetAllowedLayers(layerDef.Name)
		if err != nil {
			return nil, err
		}
		for _, destinationLayerName := range allowedLayers {
			rulesets[layerDef.Name] = make(map[string]int)
			rulesets[layerDef.Name][destinationLayerName] = 0
		}
	}
	return rulesets, nil
}

func (a *RulesetUsageAnalyser) findRulesetUsages(rulesets map[string]map[string]int) (map[string]map[string]int, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	dependencies, err := a.dependencyResolver.Resolve(astMap)
	if err != nil {
		return nil, err
	}
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		dependerLayerNames, errGet := a.layerResolver.GetLayersForReference(a.tokenResolver.Resolve(dependency.GetDepender(), astMap))
		if errGet != nil {
			return nil, errGet
		}
		for dependerLayerName := range dependerLayerNames {
			dependentLayerNames, errGetDependent := a.layerResolver.GetLayersForReference(a.tokenResolver.Resolve(dependency.GetDependent(), astMap))
			if errGetDependent != nil {
				return nil, errGetDependent
			}
			for dependentLayerName := range dependentLayerNames {
				if _, ok1 := rulesets[dependerLayerName]; ok1 {
					if _, ok2 := rulesets[dependerLayerName][dependentLayerName]; ok2 {
						rulesets[dependerLayerName][dependentLayerName]++
					}
				}
			}
		}
	}
	return rulesets, nil
}
