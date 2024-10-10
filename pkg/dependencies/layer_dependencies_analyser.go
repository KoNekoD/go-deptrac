package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type LayerDependenciesAnalyser struct {
	astMapExtractor    *ast_map.AstMapExtractor
	tokenResolver      *tokens.TokenResolver
	dependencyResolver *DependencyResolver
	layerResolver      layers.LayerResolverInterface
}

func NewLayerDependenciesAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	tokenResolver *tokens.TokenResolver,
	dependencyResolver *DependencyResolver,
	layerResolver layers.LayerResolverInterface,
) *LayerDependenciesAnalyser {
	return &LayerDependenciesAnalyser{
		astMapExtractor:    astMapExtractor,
		tokenResolver:      tokenResolver,
		dependencyResolver: dependencyResolver,
		layerResolver:      layerResolver,
	}
}

func (a *LayerDependenciesAnalyser) GetDependencies(layer string, targetLayer *string) (map[string][]*rules.Uncovered, error) {
	uncoveredResult := make(map[string][]*rules.Uncovered)
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
					uncoveredResult[dependentLayerName] = make([]*rules.Uncovered, 0)
				}
				uncoveredResult[dependentLayerName] = append(uncoveredResult[dependentLayerName], rules.NewUncovered(dependency, dependentLayerName))
			}
		}
	}

	return uncoveredResult, nil
}
