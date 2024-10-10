package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/configs"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"slices"
)

type TokenInLayerAnalyser struct {
	tokenTypes      []TokenType
	config          *configs.AnalyserConfig
	astMapExtractor *ast_map.AstMapExtractor
	tokenResolver   *TokenResolver
	layerResolver   layers.LayerResolverInterface
}

func NewTokenInLayerAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	tokenResolver *TokenResolver,
	layerResolver layers.LayerResolverInterface,
	config *configs.AnalyserConfig,
) *TokenInLayerAnalyser {
	analyser := &TokenInLayerAnalyser{
		tokenTypes:      make([]TokenType, 0),
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
		config:          config,
	}

	for _, configType := range config.Types {
		newTokenType := NewTokenTypeTryFromEmitterType(configType)
		if newTokenType == nil {
			continue
		}

		analyser.tokenTypes = append(analyser.tokenTypes, *newTokenType)
	}

	return analyser
}

func (a *TokenInLayerAnalyser) FindTokensInLayer(layer string) (map[string]TokenType, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	matchingTokens := make(map[string]TokenType)

	if slices.Contains(a.tokenTypes, TokenTypeClassLike) {
		for _, classReference := range astMap.GetClassLikeReferences() {
			classToken := a.tokenResolver.Resolve(classReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(classToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[classToken.GetToken().ToString()] = TokenTypeClassLike
			}
		}
	}

	if slices.Contains(a.tokenTypes, TokenTypeFunction) {
		for _, functionReference := range astMap.GetFunctionReferences() {
			functionToken := a.tokenResolver.Resolve(functionReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(functionToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[functionToken.GetToken().ToString()] = TokenTypeFunction
			}
		}
	}

	if slices.Contains(a.tokenTypes, TokenTypeFile) {
		for _, fileReference := range astMap.GetFileReferences() {
			fileToken := a.tokenResolver.Resolve(fileReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(fileToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[fileToken.GetToken().ToString()] = TokenTypeFile
			}
		}
	}

	return matchingTokens, nil
}
