package analyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/dependency_resolver"
	Layer2 "github.com/KoNekoD/go-deptrac/pkg/src/core/layer/layer_resolver_interface"
)

type RulesetUsageAnalyser struct {
	layerProvider      *LayerProvider.LayerProvider
	layerResolver      Layer2.LayerResolverInterface
	astMapExtractor    *ast.AstMapExtractor
	dependencyResolver *dependency_resolver.DependencyResolver
	tokenResolver      *dependency.TokenResolver
	layers             []*Layer.Layer
}

func NewRulesetUsageAnalyser(
	layerProvider *LayerProvider.LayerProvider,
	layerResolver Layer2.LayerResolverInterface,
	astMapExtractor *ast.AstMapExtractor,
	dependencyResolver *dependency_resolver.DependencyResolver,
	tokenResolver *dependency.TokenResolver,
	layers []*Layer.Layer,
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
