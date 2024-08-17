package UsesCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type UsesCollector struct {
	astMapExtractor *AstMapExtractor.AstMapExtractor
	astMap          *AstMap.AstMap
}

func NewUsesCollector(astMapExtractor *AstMapExtractor.AstMapExtractor) (*UsesCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &UsesCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (u *UsesCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.ClassLikeReference); !ok {
		return false, nil
	}

	traitName, err := u.getTraitName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range u.astMap.GetClassInherits(reference.GetToken().(*AstMap.ClassLikeToken)) {
		if AstMap.AstInheritTypeUses == inherit.Type && inherit.ClassLikeName.Equals(traitName) {
			return true, nil
		}
	}

	return false, nil
}

func (u *UsesCollector) getTraitName(config map[string]interface{}) (*AstMap.ClassLikeToken, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("UsesCollector needs the trait name as a string.")
	}

	return AstMap.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
