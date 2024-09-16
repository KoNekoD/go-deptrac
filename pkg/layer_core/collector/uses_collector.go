package collector

import (
	astContract "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type UsesCollector struct {
	astMapExtractor *ast_core.AstMapExtractor
	astMap          *ast_map2.AstMap
}

func NewUsesCollector(astMapExtractor *ast_core.AstMapExtractor) (*UsesCollector, error) {
	astMap, err := astMapExtractor.Extract()
	if err != nil {
		return nil, err
	}
	return &UsesCollector{
		astMapExtractor: astMapExtractor,
		astMap:          astMap,
	}, nil
}

func (u *UsesCollector) Satisfy(config map[string]interface{}, reference astContract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map2.ClassLikeReference); !ok {
		return false, nil
	}

	traitName, err := u.getTraitName(config)
	if err != nil {
		return false, err
	}

	for _, inherit := range u.astMap.GetClassInherits(reference.GetToken().(*ast_map2.ClassLikeToken)) {
		if ast_map2.AstInheritTypeUses == inherit.Type && inherit.ClassLikeName.Equals(traitName) {
			return true, nil
		}
	}

	return false, nil
}

func (u *UsesCollector) getTraitName(config map[string]interface{}) (*ast_map2.ClassLikeToken, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return nil, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("UsesCollector needs the trait name as a string.")
	}

	return ast_map2.NewClassLikeTokenFromFQCN(config["value"].(string)), nil
}
