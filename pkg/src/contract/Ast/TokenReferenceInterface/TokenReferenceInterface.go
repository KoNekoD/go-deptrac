package TokenReferenceInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
)

// TokenReferenceInterface - Represents the AST-TokenInterface and its location.
type TokenReferenceInterface interface {
	GetFilepath() *string
	GetToken() TokenInterface.TokenInterface
}
