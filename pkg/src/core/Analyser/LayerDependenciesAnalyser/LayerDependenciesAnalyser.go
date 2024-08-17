package LayerDependenciesAnalyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/TokenResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/LayerResolverInterface"
)

type LayerDependenciesAnalyser struct {
	astMapExtractor    *AstMapExtractor.AstMapExtractor
	tokenResolver      *TokenResolver.TokenResolver
	dependencyResolver *DependencyResolver.DependencyResolver
	layerResolver      LayerResolverInterface.LayerResolverInterface
}

func NewLayerDependenciesAnalyser(
	astMapExtractor *AstMapExtractor.AstMapExtractor,
	tokenResolver *TokenResolver.TokenResolver,
	dependencyResolver *DependencyResolver.DependencyResolver,
	layerResolver LayerResolverInterface.LayerResolverInterface,
) *LayerDependenciesAnalyser {
	return &LayerDependenciesAnalyser{
		astMapExtractor:    astMapExtractor,
		tokenResolver:      tokenResolver,
		dependencyResolver: dependencyResolver,
		layerResolver:      layerResolver,
	}
}

func (a *LayerDependenciesAnalyser) GetDependencies(layer string, targetLayer *string) (map[string][]*Uncovered.Uncovered, error) {
	result := make(map[string][]*Uncovered.Uncovered)
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
				if _, ok := result[dependentLayerName]; !ok {
					result[dependentLayerName] = make([]*Uncovered.Uncovered, 0)
				}
				result[dependentLayerName] = append(result[dependentLayerName], Uncovered.NewUncovered(dependency, dependentLayerName))
			}
		}
	}

	// TODO: Add
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

	return result, nil
}
