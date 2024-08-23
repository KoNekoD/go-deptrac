package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
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

func (v *VariableReference) GetToken() TokenInterface.TokenInterface {
	return v.tokenName
}
