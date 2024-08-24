package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
)

type VariableReference struct {
	tokenName *SuperGlobalToken
}

func NewVariableReference(tokenName *SuperGlobalToken) *VariableReference {
	return &VariableReference{tokenName: tokenName}
}

func (v *VariableReference) GetFilepath() *string {
	return nil
}

func (v *VariableReference) GetToken() ast.TokenInterface {
	return v.tokenName
}
