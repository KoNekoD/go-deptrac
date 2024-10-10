package analysers

import (
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
)

type LayerDependenciesAnalyser struct {
	astMapExtractor    *ast_map.AstMapExtractor
	tokenResolver      *services.TokenResolver
	dependencyResolver *pkg.DependencyResolver
	layerResolver      layers_resolvers.LayerResolverInterface
}

func NewLayerDependenciesAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	tokenResolver *services.TokenResolver,
	dependencyResolver *pkg.DependencyResolver,
	layerResolver layers_resolvers.LayerResolverInterface,
) *LayerDependenciesAnalyser {
	return &LayerDependenciesAnalyser{
		astMapExtractor:    astMapExtractor,
		tokenResolver:      tokenResolver,
		dependencyResolver: dependencyResolver,
		layerResolver:      layerResolver,
	}
}

func (a *LayerDependenciesAnalyser) GetDependencies(layer string, targetLayer *string) (map[string][]*violations_rules.Uncovered, error) {
	uncoveredResult := make(map[string][]*violations_rules.Uncovered)
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
					uncoveredResult[dependentLayerName] = make([]*violations_rules.Uncovered, 0)
				}
				uncoveredResult[dependentLayerName] = append(uncoveredResult[dependentLayerName], violations_rules.NewUncovered(dependency, dependentLayerName))
			}
		}
	}

	return uncoveredResult, nil
}
