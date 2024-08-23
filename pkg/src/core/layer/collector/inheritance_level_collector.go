package collector

import (
	astContract "github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type InheritanceLevelCollector struct {
	astMapExtractor *ast.AstMapExtractor
	astMap          *ast_map.AstMap
}

func NewInheritanceLevelCollector(astMapExtractor *ast.AstMapExtractor) (*InheritanceLevelCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritanceLevelCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritanceLevelCollector) Satisfy(config map[string]interface{}, reference astContract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map.ClassLikeReference); !ok {
		return false, nil
	}

	classInherits := c.astMap.GetClassInherits(reference.GetToken().(*ast_map.ClassLikeToken))

	if !util.MapKeyExists(config, "value") || util.MapKeyIsInt(config, "value") {
		return false, layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("InheritanceLevelCollector needs inheritance depth as int.")
	}

	depth := config["value"].(int)

	for _, classInherit := range classInherits {
		if len(classInherit.GetPath()) > depth {
			return true, nil
		}
	}

	return false, nil
}
