package LayerForTokenAnalyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/TokenType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/TokenResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/LayerResolverInterface"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"strings"
)

type LayerForTokenAnalyser struct {
	astMapExtractor *AstMapExtractor.AstMapExtractor
	tokenResolver   *TokenResolver.TokenResolver
	layerResolver   LayerResolverInterface.LayerResolverInterface
}

func NewLayerForTokenAnalyser(
	astMapExtractor *AstMapExtractor.AstMapExtractor,
	tokenResolver *TokenResolver.TokenResolver,
	layerResolver LayerResolverInterface.LayerResolverInterface,
) *LayerForTokenAnalyser {
	return &LayerForTokenAnalyser{
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
	}
}

func (a *LayerForTokenAnalyser) FindLayerForToken(tokenName string, tokenType TokenType.TokenType) (map[string][]string, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}

	// TODO: Add
	// } catch (UnrecognizedTokenException $e) {
	//     throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::unrecognizedToken($e);
	// } catch (InvalidLayerDefinitionException $e) {
	//     throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::invalidLayerDefinition($e);
	// } catch (InvalidCollectorDefinitionException $e) {
	//     throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::invalidCollectorDefinition($e);
	// } catch (AstException $e) {
	//     throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::failedAstParsing($e);
	// } catch (CouldNotParseFileException $e) {
	//     throw \Qossmic\Deptrac\Core\SetAnalyser\AnalyserException::couldNotParseFile($e);
	// }

	switch tokenType {
	case TokenType.TokenTypeClassLike:
		return a.findLayersForReferences(astMap.GetClassLikeReferences(), tokenName, astMap)
	case TokenType.TokenTypeFunction:
		return a.findLayersForReferences(astMap.GetFunctionReferences(), tokenName, astMap)
	case TokenType.TokenTypeFile:
		return a.findLayersForReferences(astMap.GetFileReferences(), tokenName, astMap)
	default:
		return nil, errors.New("Invalid token type")
	}
}

func (a *LayerForTokenAnalyser) findLayersForReferences(referencesAny any, tokenName string, astMap *AstMap.AstMap) (map[string][]string, error) {
	references := referencesAny.([]TokenReferenceInterface.TokenReferenceInterface)
	if len(references) == 0 {
		return make(map[string][]string), nil
	}

	layersForReference := make(map[string][]string)

	for _, reference := range references {
		if !strings.Contains(reference.GetToken().ToString(), tokenName) {
			continue
		}
		token := a.tokenResolver.Resolve(reference.GetToken(), astMap)
		gotLayers, err := a.layerResolver.GetLayersForReference(token)

		if err != nil {
			return nil, err
		}

		layersForReference[reference.GetToken().ToString()] = maps.Keys(gotLayers)
	}

	return layersForReference, nil
}
