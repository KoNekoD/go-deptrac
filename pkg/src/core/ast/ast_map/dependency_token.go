package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type DependencyToken struct {
	Token   ast.TokenInterface
	Context *ast.DependencyContext
}

func NewDependencyToken(token ast.TokenInterface, context *ast.DependencyContext) *DependencyToken {
	return &DependencyToken{Token: token, Context: context}
}
