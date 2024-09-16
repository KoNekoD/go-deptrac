package ast_map

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

type DependencyToken struct {
	Token   ast_contract2.TokenInterface
	Context *ast_contract2.DependencyContext
}

func NewDependencyToken(token ast_contract2.TokenInterface, context *ast_contract2.DependencyContext) *DependencyToken {
	return &DependencyToken{Token: token, Context: context}
}
