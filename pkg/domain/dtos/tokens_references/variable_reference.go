package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type VariableReference struct {
	tokenName *enums.SuperGlobalToken
}

func NewVariableReference(tokenName *enums.SuperGlobalToken) *VariableReference {
	return &VariableReference{tokenName: tokenName}
}

func (v *VariableReference) GetFilepath() *string {
	return nil
}

func (v *VariableReference) GetToken() tokens.TokenInterface {
	return v.tokenName
}
