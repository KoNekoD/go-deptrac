package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

type InheritsCollector struct {
	astMapExtractor *ast_map.AstMapExtractor
	astMap          *ast_map.AstMap
}

func NewInheritsCollector(astMapExtractor *ast_map.AstMapExtractor) (*InheritsCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritsCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritsCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*tokens_references2.ClassLikeReference); !ok {
		return false, nil
	}

	classLikeName, err := c.getClassLikeName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range c.astMap.GetClassInherits(reference.GetToken().(*tokens2.ClassLikeToken)) {
		if inherit.ClassLikeName.Equals(classLikeName) {
			return true, nil
		}
	}

	return false, nil
}

func (c *InheritsCollector) getClassLikeName(config map[string]interface{}) (*tokens2.ClassLikeToken, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return nil, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("InheritsCollector needs the interface, trait or class name as a string.")
	}

	return tokens2.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
