package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type TokenResolver struct{}

func NewTokenResolver() *TokenResolver {
	return &TokenResolver{}
}

func (r *TokenResolver) Resolve(token tokens.TokenInterface, astMap *ast_maps.AstMap) tokens_references.TokenReferenceInterface {
	switch v := token.(type) {
	case *tokens.ClassLikeToken:
		return astMap.GetClassReferenceForToken(v)
	case *tokens.FunctionToken:
		return astMap.GetFunctionReferenceForToken(v)
	case *tokens.FileToken:
		return astMap.GetFileReferenceForToken(v)
	default:
		panic("Unrecognized TokenInterface")
	}
}
