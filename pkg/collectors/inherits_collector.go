package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
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

func (c *InheritsCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*references.ClassLikeReference); !ok {
		return false, nil
	}

	classLikeName, err := c.getClassLikeName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range c.astMap.GetClassInherits(reference.GetToken().(*tokens.ClassLikeToken)) {
		if inherit.ClassLikeName.Equals(classLikeName) {
			return true, nil
		}
	}

	return false, nil
}

func (c *InheritsCollector) getClassLikeName(config map[string]interface{}) (*tokens.ClassLikeToken, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return nil, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("InheritsCollector needs the interface, trait or class name as a string.")
	}

	return tokens.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
