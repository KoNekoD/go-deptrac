package collector

import (
	astContract "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type InheritsCollector struct {
	astMapExtractor *ast_core.AstMapExtractor
	astMap          *ast_map2.AstMap
}

func NewInheritsCollector(astMapExtractor *ast_core.AstMapExtractor) (*InheritsCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritsCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritsCollector) Satisfy(config map[string]interface{}, reference astContract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map2.ClassLikeReference); !ok {
		return false, nil
	}

	classLikeName, err := c.getClassLikeName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range c.astMap.GetClassInherits(reference.GetToken().(*ast_map2.ClassLikeToken)) {
		if inherit.ClassLikeName.Equals(classLikeName) {
			return true, nil
		}
	}

	return false, nil
}

func (c *InheritsCollector) getClassLikeName(config map[string]interface{}) (*ast_map2.ClassLikeToken, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return nil, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("InheritsCollector needs the interface, trait or class name as a string.")
	}

	return ast_map2.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
