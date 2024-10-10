package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"strings"
)

type LayerForTokenAnalyser struct {
	astMapExtractor *ast_map.AstMapExtractor
	tokenResolver   *TokenResolver
	layerResolver   layers.LayerResolverInterface
}

func NewLayerForTokenAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	tokenResolver *TokenResolver,
	layerResolver layers.LayerResolverInterface,
) *LayerForTokenAnalyser {
	return &LayerForTokenAnalyser{
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
	}
}

func (a *LayerForTokenAnalyser) FindLayerForToken(tokenName string, tokenType enums.TokenType) (map[string][]string, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}

	switch tokenType {
	case enums.TokenTypeClassLike:
		return a.findLayersForReferences(astMap.GetClassLikeReferences(), tokenName, astMap)
	case enums.TokenTypeFunction:
		return a.findLayersForReferences(astMap.GetFunctionReferences(), tokenName, astMap)
	case enums.TokenTypeFile:
		return a.findLayersForReferences(astMap.GetFileReferences(), tokenName, astMap)
	default:
		return nil, errors.New("Invalid token type")
	}
}

func (a *LayerForTokenAnalyser) findLayersForReferences(referencesAny any, tokenName string, astMap *ast_map2.AstMap) (map[string][]string, error) {
	references := referencesAny.([]tokens2.TokenReferenceInterface)
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
