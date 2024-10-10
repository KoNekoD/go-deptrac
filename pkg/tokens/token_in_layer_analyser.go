package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"slices"
)

type TokenInLayerAnalyser struct {
	tokenTypes      []enums.TokenType
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
		tokenTypes:      make([]enums.TokenType, 0),
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
		config:          config,
	}

	for _, configType := range config.Types {
		newTokenType := enums.NewTokenTypeTryFromEmitterType(configType)
		if newTokenType == nil {
			continue
		}

		analyser.tokenTypes = append(analyser.tokenTypes, *newTokenType)
	}

	return analyser
}

func (a *TokenInLayerAnalyser) FindTokensInLayer(layer string) (map[string]enums.TokenType, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	matchingTokens := make(map[string]enums.TokenType)

	if slices.Contains(a.tokenTypes, enums.TokenTypeClassLike) {
		for _, classReference := range astMap.GetClassLikeReferences() {
			classToken := a.tokenResolver.Resolve(classReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(classToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[classToken.GetToken().ToString()] = enums.TokenTypeClassLike
			}
		}
	}

	if slices.Contains(a.tokenTypes, enums.TokenTypeFunction) {
		for _, functionReference := range astMap.GetFunctionReferences() {
			functionToken := a.tokenResolver.Resolve(functionReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(functionToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[functionToken.GetToken().ToString()] = enums.TokenTypeFunction
			}
		}
	}

	if slices.Contains(a.tokenTypes, enums.TokenTypeFile) {
		for _, fileReference := range astMap.GetFileReferences() {
			fileToken := a.tokenResolver.Resolve(fileReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(fileToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[fileToken.GetToken().ToString()] = enums.TokenTypeFile
			}
		}
	}

	return matchingTokens, nil
}
