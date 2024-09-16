package dependency_core

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token ast_contract2.TokenInterface, astMap *ast_map.AstMap) ast_contract2.TokenReferenceInterface {
	switch v := token.(type) {
	case *ast_map.ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *ast_map.FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *ast_map.SuperGlobalToken:
		return ast_map.NewVariableReference(v)
	case *ast_map.FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}
