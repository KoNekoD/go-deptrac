package InheritsCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type InheritsCollector struct {
	astMapExtractor *AstMapExtractor.AstMapExtractor
	astMap          *AstMap.AstMap
}

func NewInheritsCollector(astMapExtractor *AstMapExtractor.AstMapExtractor) (*InheritsCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritsCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritsCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.ClassLikeReference); !ok {
		return false, nil
	}

	classLikeName, err := c.getClassLikeName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range c.astMap.GetClassInherits(reference.GetToken().(*AstMap.ClassLikeToken)) {
		if inherit.ClassLikeName.Equals(classLikeName) {
			return true, nil
		}
	}

	return false, nil
}

func (c *InheritsCollector) getClassLikeName(config map[string]interface{}) (*AstMap.ClassLikeToken, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("InheritsCollector needs the interface, trait or class name as a string.")
	}

	return AstMap.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
