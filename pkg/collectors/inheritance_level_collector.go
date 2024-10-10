package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type InheritanceLevelCollector struct {
	astMapExtractor *ast_map.AstMapExtractor
	astMap          *ast_map.AstMap
}

func NewInheritanceLevelCollector(astMapExtractor *ast_map.AstMapExtractor) (*InheritanceLevelCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &InheritanceLevelCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (c *InheritanceLevelCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*references.ClassLikeReference); !ok {
		return false, nil
	}

	classInherits := c.astMap.GetClassInherits(reference.GetToken().(*tokens.ClassLikeToken))

	if !utils.MapKeyExists(config, "value") || utils.MapKeyIsInt(config, "value") {
		return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("InheritanceLevelCollector needs inheritance depth as int.")
	}

	depth := config["value"].(int)

	for _, classInherit := range classInherits {
		if len(classInherit.GetPath()) > depth {
			return true, nil
		}
	}

	return false, nil
}
