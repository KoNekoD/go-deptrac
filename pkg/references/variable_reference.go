package references

import "github.com/KoNekoD/go-deptrac/pkg/tokens"

type VariableReference struct {
	tokenName *tokens.SuperGlobalToken
}

func NewVariableReference(tokenName *tokens.SuperGlobalToken) *VariableReference {
	return &VariableReference{tokenName: tokenName}
}

func (v *VariableReference) GetFilepath() *string {
	return nil
}

func (v *VariableReference) GetToken() tokens.TokenInterface {
	return v.tokenName
}
