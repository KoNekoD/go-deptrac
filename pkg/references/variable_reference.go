package references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
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
