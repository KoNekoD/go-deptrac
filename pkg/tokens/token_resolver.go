package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/references"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token TokenInterface, astMap *ast_map.AstMap) TokenReferenceInterface {
	switch v := token.(type) {
	case *ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *SuperGlobalToken:
		return references.NewVariableReference(v)
	case *FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}
