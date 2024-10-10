package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

type UsesCollector struct {
	astMapExtractor *ast_map.AstMapExtractor
	astMap          *ast_map2.AstMap
}

func NewUsesCollector(astMapExtractor *ast_map.AstMapExtractor) (*UsesCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &UsesCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (u *UsesCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*tokens_references2.ClassLikeReference); !ok {
		return false, nil
	}

	traitName, err := u.getTraitName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range u.astMap.GetClassInherits(reference.GetToken().(*tokens2.ClassLikeToken)) {
		if enums.AstInheritTypeUses == inherit.Type && inherit.ClassLikeName.Equals(traitName) {
			return true, nil
		}
	}

	return false, nil
}

func (u *UsesCollector) getTraitName(config map[string]interface{}) (*tokens2.ClassLikeToken, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return nil, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("UsesCollector needs the trait name as a string.")
	}

	return tokens2.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
