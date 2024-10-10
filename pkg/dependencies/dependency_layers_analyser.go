package dependencies

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type DependencyLayersAnalyser struct {
	astMapExtractor    *ast_map.AstMapExtractor
	dependencyResolver *DependencyResolver
	tokenResolver      *tokens.TokenResolver
	layerResolver      layers.LayerResolverInterface
	eventDispatcher    events.EventDispatcherInterface
}

func NewDependencyLayersAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	dependencyResolver *DependencyResolver,
	tokenResolver *tokens.TokenResolver,
	layerResolver layers.LayerResolverInterface,
	eventDispatcher events.EventDispatcherInterface) *DependencyLayersAnalyser {
	return &DependencyLayersAnalyser{
		astMapExtractor:    astMapExtractor,
		dependencyResolver: dependencyResolver,
		tokenResolver:      tokenResolver,
		layerResolver:      layerResolver,
		eventDispatcher:    eventDispatcher,
	}
}

func (a *DependencyLayersAnalyser) Analyse() (*rules.AnalysisResult, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	dependencies, err := a.dependencyResolver.Resolve(astMap)
	if err != nil {
		return nil, err
	}
	analysisResult := rules.NewAnalysisResult()
	warnings := make(map[string]*tokens.Warning)
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		depender := dependency.GetDepender()
		dependerRef := a.tokenResolver.Resolve(depender, astMap)

		if v, ok55 := dependerRef.(*references.FunctionReference); ok55 {
			t := v.GetToken()
			if tt, ok66 := t.(*tokens.FunctionToken); ok66 {
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
			warnings[depender.ToString()] = tokens.NewWarningTokenIsInMoreThanOneLayer(depender.ToString(), dependerLayers)
		}

		dependent := dependency.GetDependent()
		dependentRef := a.tokenResolver.Resolve(dependent, astMap)

		dependentLayers, err := a.layerResolver.GetLayersForReference(dependentRef)
		if err != nil {
			return nil, err
		}

		for _, dependerLayer := range dependerLayers {
			event := events.NewProcessEvent(dependency, dependerRef, dependerLayer, dependentRef, dependentLayers, analysisResult)
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

	event := events.NewPostProcessEvent(analysisResult)
	errDispatch := a.eventDispatcher.DispatchEvent(event)
	if errDispatch != nil {
		return nil, errDispatch
	}

	return event.GetResult(), nil
}
