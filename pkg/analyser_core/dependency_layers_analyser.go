package analyser_core

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher/event_dispatcher_interface"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract"
)

type DependencyLayersAnalyser struct {
	astMapExtractor    *ast_core.AstMapExtractor
	dependencyResolver *dependency_resolver.DependencyResolver
	tokenResolver      *dependency_core.TokenResolver
	layerResolver      layer_resolver_interface.LayerResolverInterface
	eventDispatcher    event_dispatcher_interface.EventDispatcherInterface
}

func NewDependencyLayersAnalyser(
	astMapExtractor *ast_core.AstMapExtractor,
	dependencyResolver *dependency_resolver.DependencyResolver,
	tokenResolver *dependency_core.TokenResolver,
	layerResolver layer_resolver_interface.LayerResolverInterface,
	eventDispatcher event_dispatcher_interface.EventDispatcherInterface) *DependencyLayersAnalyser {
	return &DependencyLayersAnalyser{
		astMapExtractor:    astMapExtractor,
		dependencyResolver: dependencyResolver,
		tokenResolver:      tokenResolver,
		layerResolver:      layerResolver,
		eventDispatcher:    eventDispatcher,
	}
}

func (a *DependencyLayersAnalyser) Analyse() (*analysis_result.AnalysisResult, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	dependencies, err := a.dependencyResolver.Resolve(astMap)
	if err != nil {
		return nil, err
	}
	analysisResult := analysis_result.NewAnalysisResult()
	warnings := make(map[string]*result_contract.Warning)
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		depender := dependency.GetDepender()
		dependerRef := a.tokenResolver.Resolve(depender, astMap)

		if v, ok55 := dependerRef.(*ast_map2.FunctionReference); ok55 {
			t := v.GetToken()
			if tt, ok66 := t.(*ast_map2.FunctionToken); ok66 {
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
			warnings[depender.ToString()] = result_contract.NewWarningTokenIsInMoreThanOneLayer(depender.ToString(), dependerLayers)
		}

		dependent := dependency.GetDependent()
		dependentRef := a.tokenResolver.Resolve(dependent, astMap)

		dependentLayers, err := a.layerResolver.GetLayersForReference(dependentRef)
		if err != nil {
			return nil, err
		}

		for _, dependerLayer := range dependerLayers {
			event := process_event.NewProcessEvent(dependency, dependerRef, dependerLayer, dependentRef, dependentLayers, analysisResult)
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

	event := post_process_event.NewPostProcessEvent(analysisResult)
	errDispatch := a.eventDispatcher.DispatchEvent(event)
	if errDispatch != nil {
		return nil, errDispatch
	}

	return event.GetResult(), nil
}
