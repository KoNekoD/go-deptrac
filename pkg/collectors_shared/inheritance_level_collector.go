package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
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

func (c *InheritanceLevelCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*tokens_references2.ClassLikeReference); !ok {
		return false, nil
	}

	classInherits := c.astMap.GetClassInherits(reference.GetToken().(*tokens2.ClassLikeToken))

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
