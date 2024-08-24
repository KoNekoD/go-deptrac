package analyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/core/layer/layer_resolver_interface"
)

type LayerDependenciesAnalyser struct {
	astMapExtractor    *ast.AstMapExtractor
	tokenResolver      *dependency.TokenResolver
	dependencyResolver *dependency_resolver.DependencyResolver
	layerResolver      layer_resolver_interface.LayerResolverInterface
}

func NewLayerDependenciesAnalyser(
	astMapExtractor *ast.AstMapExtractor,
	tokenResolver *dependency.TokenResolver,
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

func (a *LayerDependenciesAnalyser) GetDependencies(layer string, targetLayer *string) (map[string][]*result.Uncovered, error) {
	uncoveredResult := make(map[string][]*result.Uncovered)
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
					uncoveredResult[dependentLayerName] = make([]*result.Uncovered, 0)
				}
				uncoveredResult[dependentLayerName] = append(uncoveredResult[dependentLayerName], result.NewUncovered(dependency, dependentLayerName))
			}
		}
	}

	return uncoveredResult, nil
}
