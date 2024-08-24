package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	AstMap2 "github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token ast.TokenInterface, astMap *AstMap2.AstMap) ast.TokenReferenceInterface {
	switch v := token.(type) {
	case *AstMap2.ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *AstMap2.FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *AstMap2.SuperGlobalToken:
		return AstMap2.NewVariableReference(v)
	case *AstMap2.FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}
