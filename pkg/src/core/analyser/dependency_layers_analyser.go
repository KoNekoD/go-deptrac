package analyser

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Warning"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
)

type DependencyLayersAnalyser struct {
	astMapExtractor    *ast.AstMapExtractor
	dependencyResolver *dependency_resolver.DependencyResolver
	tokenResolver      *dependency.TokenResolver
	layerResolver      layer_resolver_interface.LayerResolverInterface
	eventDispatcher    util.EventDispatcherInterface
}

func NewDependencyLayersAnalyser(
	astMapExtractor *ast.AstMapExtractor,
	dependencyResolver *dependency_resolver.DependencyResolver,
	tokenResolver *dependency.TokenResolver,
	layerResolver layer_resolver_interface.LayerResolverInterface,
	eventDispatcher util.EventDispatcherInterface) *DependencyLayersAnalyser {
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
	result := analysis_result.NewAnalysisResult()
	warnings := make(map[string]*Warning.Warning)
	for _, dependency := range dependencies.GetDependenciesAndInheritDependencies() {
		depender := dependency.GetDepender()
		dependerRef := a.tokenResolver.Resolve(depender, astMap)

		if v, ok55 := dependerRef.(*ast_map.FunctionReference); ok55 {
			t := v.GetToken()
			if tt, ok66 := t.(*ast_map.FunctionToken); ok66 {
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
			warnings[depender.ToString()] = Warning.NewWarningTokenIsInMoreThanOneLayer(depender.ToString(), dependerLayers)
		}

		dependent := dependency.GetDependent()
		dependentRef := a.tokenResolver.Resolve(dependent, astMap)

		dependentLayers, err := a.layerResolver.GetLayersForReference(dependentRef)
		if err != nil {
			return nil, err
		}

		for _, dependerLayer := range dependerLayers {
			event := process_event.NewProcessEvent(dependency, dependerRef, dependerLayer, dependentRef, dependentLayers, result)
			err := a.eventDispatcher.DispatchEvent(event)
			if err != nil {
				return nil, err
			}
			result = event.GetResult()
		}
	}

	for _, warning := range warnings {
		result.AddWarning(warning)
	}

	event := post_process_event.NewPostProcessEvent(result)
	errDispatch := a.eventDispatcher.DispatchEvent(event)
	if errDispatch != nil {
		return nil, errDispatch
	}

	//TODO: Implement
	//} catch (InvalidEmitterConfigurationException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::invalidEmitterConfiguration($e);
	//} catch (UnrecognizedTokenException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::unrecognizedToken($e);
	//} catch (InvalidLayerDefinitionException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::invalidLayerDefinition($e);
	//} catch (InvalidCollectorDefinitionException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::invalidCollectorDefinition($e);
	//} catch (AstException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::failedAstParsing($e);
	//} catch (CouldNotParseFileException $e) {
	//    throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::couldNotParseFile($e);
	//}

	return event.GetResult(), nil
}