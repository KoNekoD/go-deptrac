package analyser_core

import (
	astContract "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core/layer_resolver_interface"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"strings"
)

type LayerForTokenAnalyser struct {
	astMapExtractor *ast_core.AstMapExtractor
	tokenResolver   *dependency_core.TokenResolver
	layerResolver   layer_resolver_interface.LayerResolverInterface
}

func NewLayerForTokenAnalyser(
	astMapExtractor *ast_core.AstMapExtractor,
	tokenResolver *dependency_core.TokenResolver,
	layerResolver layer_resolver_interface.LayerResolverInterface,
) *LayerForTokenAnalyser {
	return &LayerForTokenAnalyser{
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
	}
}

func (a *LayerForTokenAnalyser) FindLayerForToken(tokenName string, tokenType TokenType) (map[string][]string, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}

	switch tokenType {
	case TokenTypeClassLike:
		return a.findLayersForReferences(astMap.GetClassLikeReferences(), tokenName, astMap)
	case TokenTypeFunction:
		return a.findLayersForReferences(astMap.GetFunctionReferences(), tokenName, astMap)
	case TokenTypeFile:
		return a.findLayersForReferences(astMap.GetFileReferences(), tokenName, astMap)
	default:
		return nil, errors.New("Invalid token type")
	}
}

func (a *LayerForTokenAnalyser) findLayersForReferences(referencesAny any, tokenName string, astMap *ast_map.AstMap) (map[string][]string, error) {
	references := referencesAny.([]astContract.TokenReferenceInterface)
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
