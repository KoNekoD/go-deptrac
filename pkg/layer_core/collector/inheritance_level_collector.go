package collector

import (
	astContract "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type InheritanceLevelCollector struct {
	astMapExtractor *ast_core.AstMapExtractor
	astMap          *ast_map2.AstMap
}

func NewInheritanceLevelCollector(astMapExtractor *ast_core.AstMapExtractor) (*InheritanceLevelCollector, error) {
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
	if _, ok := reference.(*ast_map2.ClassLikeReference); !ok {
		return false, nil
	}

	classInherits := c.astMap.GetClassInherits(reference.GetToken().(*ast_map2.ClassLikeToken))

	if !util.MapKeyExists(config, "value") || util.MapKeyIsInt(config, "value") {
		return false, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("InheritanceLevelCollector needs inheritance depth as int.")
	}

	depth := config["value"].(int)

	for _, classInherit := range classInherits {
		if len(classInherit.GetPath()) > depth {
			return true, nil
		}
	}

	return false, nil
}
