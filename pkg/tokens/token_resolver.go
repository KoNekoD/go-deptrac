package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	tokens3 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token tokens3.TokenInterface, astMap *ast_map.AstMap) tokens_references2.TokenReferenceInterface {
	switch v := token.(type) {
	case *tokens3.ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *tokens3.FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *enums.SuperGlobalToken:
		return tokens_references2.NewVariableReference(v)
	case *tokens3.FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}
