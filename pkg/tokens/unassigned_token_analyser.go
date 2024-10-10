package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"slices"
)

type UnassignedTokenAnalyser struct {
	tokenTypes      []enums.TokenType
	config          *configs.AnalyserConfig
	astMapExtractor *ast_map.AstMapExtractor
	tokenResolver   *TokenResolver
	layerResolver   layers.LayerResolverInterface
}

func NewUnassignedTokenAnalyser(
	astMapExtractor *ast_map.AstMapExtractor,
	tokenResolver *TokenResolver,
	layerResolver layers.LayerResolverInterface,
	config *configs.AnalyserConfig,
) *UnassignedTokenAnalyser {
	analyser := &UnassignedTokenAnalyser{
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

func (a *UnassignedTokenAnalyser) FindUnassignedTokens() ([]string, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	unassignedTokens := make([]string, 0)

	if slices.Contains(a.tokenTypes, enums.TokenTypeClassLike) {
		for _, classReference := range astMap.GetClassLikeReferences() {
			classToken := a.tokenResolver.Resolve(classReference.GetToken(), astMap)
			gotLayers, errGet := a.layerResolver.GetLayersForReference(classToken)
			if errGet != nil {
				return nil, errGet
			}

			if len(gotLayers) == 0 {
				unassignedTokens = append(unassignedTokens, classToken.GetToken().ToString())
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

			if len(gotLayers) == 0 {
				unassignedTokens = append(unassignedTokens, functionToken.GetToken().ToString())
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

			if len(gotLayers) == 0 {
				unassignedTokens = append(unassignedTokens, fileToken.GetToken().ToString())
			}
		}
	}

	slices.Sort(unassignedTokens)

	return unassignedTokens, nil
}
