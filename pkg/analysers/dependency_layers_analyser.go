package analysers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/issues"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	events2 "github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type DependencyLayersAnalyser struct {
	astMapExtractor    *ast_map.AstMapExtractor
	dependencyResolver *pkg.DependencyResolver
	tokenResolver      *services.TokenResolver
	layerResolver      layers_resolvers.LayerResolverInterface
	eventDispatcher    dispatchers.EventDispatcherInterface
}

func NewDependencyLayersAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	dependencyResolver *pkg.DependencyResolver,
	tokenResolver *services.TokenResolver,
	layerResolver layers_resolvers.LayerResolverInterface,
	eventDispatcher dispatchers.EventDispatcherInterface) *DependencyLayersAnalyser {
	return &DependencyLayersAnalyser{
		astMapExtractor:    astMapExtractor,
		dependencyResolver: dependencyResolver,
		tokenResolver:      tokenResolver,
		layerResolver:      layerResolver,
		eventDispatcher:    eventDispatcher,
	}
}

func (a *DependencyLayersAnalyser) Analyse() (*results.AnalysisResult, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	dependencies, err := a.dependencyResolver.Resolve(astMap)
	if err != nil {
		return nil, err
	}
	analysisResult := results.NewAnalysisResult()
	warnings := make(map[string]*issues.Warning)
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		depender := dependency.GetDepender()
		dependerRef := a.tokenResolver.Resolve(depender, astMap)

		if v, ok55 := dependerRef.(*tokens_references.FunctionReference); ok55 {
			t := v.GetToken()
			if tt, ok66 := t.(*tokens2.FunctionToken); ok66 {
				if tt.FunctionName == "ParseFile" {
					fmt.Println()
				}
			}
		}

		dependerLayersMap, err := a.layerResolver.GetLayersForReference(dependerRef)
		if err != nil {
			return nil, err
		}
		dependerLayers := make([]string, 0)
		for s := range dependerLayersMap {
			dependerLayers = append(dependerLayers, s)
		}

		_, ok := warnings[depender.ToString()]
		if !ok && len(dependerLayers) > 1 {
			warnings[depender.ToString()] = issues.NewWarningTokenIsInMoreThanOneLayer(depender.ToString(), dependerLayers)
		}

		dependent := dependency.GetDependent()
		dependentRef := a.tokenResolver.Resolve(dependent, astMap)

		dependentLayers, err := a.layerResolver.GetLayersForReference(dependentRef)
		if err != nil {
			return nil, err
		}

		for _, dependerLayer := range dependerLayers {
			event := events2.NewProcessEvent(dependency, dependerRef, dependerLayer, dependentRef, dependentLayers, analysisResult)
			err := a.eventDispatcher.DispatchEvent(event)
			if err != nil {
				return nil, err
			}
			analysisResult = event.GetResult()
		}
	}

	for _, warning := range warnings {
		analysisResult.AddWarning(warning)
	}

	event := events2.NewPostProcessEvent(analysisResult)
	errDispatch := a.eventDispatcher.DispatchEvent(event)
	if errDispatch != nil {
		return nil, errDispatch
	}

	return event.GetResult(), nil
}
