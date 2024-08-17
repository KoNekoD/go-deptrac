package TokenInLayerAnalyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/AnalyserConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/TokenType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/TokenResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/LayerResolverInterface"
	"slices"
)

type TokenInLayerAnalyser struct {
	tokenTypes      []TokenType.TokenType
	config          *AnalyserConfig.AnalyserConfig
	astMapExtractor *AstMapExtractor.AstMapExtractor
	tokenResolver   *TokenResolver.TokenResolver
	layerResolver   LayerResolverInterface.LayerResolverInterface
}

func NewTokenInLayerAnalyser(
	astMapExtractor *AstMapExtractor.AstMapExtractor,
	tokenResolver *TokenResolver.TokenResolver,
	layerResolver LayerResolverInterface.LayerResolverInterface,
	config *AnalyserConfig.AnalyserConfig,
) *TokenInLayerAnalyser {
	analyser := &TokenInLayerAnalyser{
		tokenTypes:      make([]TokenType.TokenType, 0),
		astMapExtractor: astMapExtractor,
		tokenResolver:   tokenResolver,
		layerResolver:   layerResolver,
		config:          config,
	}

	for _, configType := range config.Types {
		newTokenType := TokenType.NewTokenTypeTryFromEmitterType(configType)
		if newTokenType == nil {
			continue
		}

		analyser.tokenTypes = append(analyser.tokenTypes, *newTokenType)
	}

	return analyser
}

func (a *TokenInLayerAnalyser) FindTokensInLayer(layer string) (map[string]TokenType.TokenType, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	matchingTokens := make(map[string]TokenType.TokenType)

	if slices.Contains(a.tokenTypes, TokenType.TokenTypeClassLike) {
		for _, classReference := range astMap.GetClassLikeReferences() {
			classToken := a.tokenResolver.Resolve(classReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(classToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[classToken.GetToken().ToString()] = TokenType.TokenTypeClassLike
			}
		}
	}

	if slices.Contains(a.tokenTypes, TokenType.TokenTypeFunction) {
		for _, functionReference := range astMap.GetFunctionReferences() {
			functionToken := a.tokenResolver.Resolve(functionReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(functionToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[functionToken.GetToken().ToString()] = TokenType.TokenTypeFunction
			}
		}
	}

	if slices.Contains(a.tokenTypes, TokenType.TokenTypeFile) {
		for _, fileReference := range astMap.GetFileReferences() {
			fileToken := a.tokenResolver.Resolve(fileReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(fileToken)
			if errGet != nil {
				return nil, errGet
			}
			if _, ok := gotLayers[layer]; ok {
				matchingTokens[fileToken.GetToken().ToString()] = TokenType.TokenTypeFile
			}
		}
	}

	return matchingTokens, nil
}
