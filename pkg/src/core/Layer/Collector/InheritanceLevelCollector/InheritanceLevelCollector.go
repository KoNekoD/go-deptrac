package InheritanceLevelCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type InheritanceLevelCollector struct {
	astMapExtractor *AstMapExtractor.AstMapExtractor
	astMap          *AstMap.AstMap
}

func NewInheritanceLevelCollector(astMapExtractor *AstMapExtractor.AstMapExtractor) (*InheritanceLevelCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritanceLevelCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritanceLevelCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.ClassLikeReference); !ok {
		return false, nil
	}

	classInherits := c.astMap.GetClassInherits(reference.GetToken().(*AstMap.ClassLikeToken))

	if !util.MapKeyExists(config, "value") || util.MapKeyIsInt(config, "value") {
		return false, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("InheritanceLevelCollector needs inheritance depth as int.")
	}

	depth := config["value"].(int)

	for _, classInherit := range classInherits {
		if len(classInherit.GetPath()) > depth {
			return true, nil
		}
	}

	return false, nil
}
