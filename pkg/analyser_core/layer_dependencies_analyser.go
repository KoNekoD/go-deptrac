package analyser_core

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract"
)

type LayerDependenciesAnalyser struct {
	astMapExtractor    *ast_core.AstMapExtractor
	tokenResolver      *dependency_core.TokenResolver
	dependencyResolver *dependency_resolver.DependencyResolver
	layerResolver      layer_resolver_interface.LayerResolverInterface
}

func NewLayerDependenciesAnalyser(
	astMapExtractor *ast_core.AstMapExtractor,
	tokenResolver *dependency_core.TokenResolver,
	dependencyResolver *dependency_resolver.DependencyResolver,
	layerResolver layer_resolver_interface.LayerResolverInterface,
) *LayerDependenciesAnalyser {
	return &LayerDependenciesAnalyser{
		astMapExtractor:    astMapExtractor,
		tokenResolver:      tokenResolver,
		dependencyResolver: dependencyResolver,
		layerResolver:      layerResolver,
	}
}

func (a *LayerDependenciesAnalyser) GetDependencies(layer string, targetLayer *string) (map[string][]*result_contract.Uncovered, error) {
	uncoveredResult := make(map[string][]*result_contract.Uncovered)
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	dependencies, err := a.dependencyResolver.Resolve(astMap)
	if err != nil {
		return nil, err
	}
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		dependerLayerNames, err := a.layerResolver.GetLayersForReference(a.tokenResolver.Resolve(dependency.GetDepender(), astMap))
		if err != nil {
			return nil, err
		}
		if _, ok := dependerLayerNames[layer]; ok {
			dependentLayerNames, err := a.layerResolver.GetLayersForReference(a.tokenResolver.Resolve(dependency.GetDependent(), astMap))
			if err != nil {
				return nil, err
			}
			for dependentLayerName := range dependentLayerNames {
				if layer == dependentLayerName || targetLayer != nil && *targetLayer != dependentLayerName {
					continue
				}
				if _, ok := uncoveredResult[dependentLayerName]; !ok {
					uncoveredResult[dependentLayerName] = make([]*result_contract.Uncovered, 0)
				}
				uncoveredResult[dependentLayerName] = append(uncoveredResult[dependentLayerName], result_contract.NewUncovered(dependency, dependentLayerName))
			}
		}
	}

	return uncoveredResult, nil
}
