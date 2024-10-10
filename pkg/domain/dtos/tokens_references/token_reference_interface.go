package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

// TokenReferenceInterface - Represents the AST-TokenInterface and its location.
type TokenReferenceInterface interface {
	GetFilepath() *string
	GetToken() tokens.TokenInterface
}
