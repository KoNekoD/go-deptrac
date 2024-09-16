package analyser_core

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core/layer_resolver_interface"
	"slices"
)

type UnassignedTokenAnalyser struct {
	tokenTypes      []TokenType
	config          *config_contract.AnalyserConfig
	astMapExtractor *ast_core.AstMapExtractor
	tokenResolver   *dependency_core.TokenResolver
	layerResolver   layer_resolver_interface.LayerResolverInterface
}

func NewUnassignedTokenAnalyser(
	astMapExtractor *ast_core.AstMapExtractor,
	tokenResolver *dependency_core.TokenResolver,
	layerResolver layer_resolver_interface.LayerResolverInterface,
	config *config_contract.AnalyserConfig,
) *UnassignedTokenAnalyser {
	analyser := &UnassignedTokenAnalyser{
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

func (a *UnassignedTokenAnalyser) FindUnassignedTokens() ([]string, error) {
	astMap, err := a.astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	unassignedTokens := make([]string, 0)

	if slices.Contains(a.tokenTypes, TokenTypeClassLike) {
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

	if slices.Contains(a.tokenTypes, TokenTypeFunction) {
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

	if slices.Contains(a.tokenTypes, TokenTypeFile) {
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
