package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token tokens.TokenInterface, astMap *ast_map.AstMap) tokens_references.TokenReferenceInterface {
	switch v := token.(type) {
	case *tokens.ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *tokens.FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *enums.SuperGlobalToken:
		return tokens_references.NewVariableReference(v)
	case *tokens.FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}