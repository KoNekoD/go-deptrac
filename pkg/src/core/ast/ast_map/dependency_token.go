package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
)

type DependencyToken struct {
	Token   TokenInterface.TokenInterface
	Context *DependencyContext.DependencyContext
}

func NewDependencyToken(token TokenInterface.TokenInterface, context *DependencyContext.DependencyContext) *DependencyToken {
	return &DependencyToken{Token: token, Context: context}
}
