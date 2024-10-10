package tokens

// TokenReferenceInterface - Represents the AST-TokenInterface and its location.
type TokenReferenceInterface interface {
	GetFilepath() *string
	GetToken() TokenInterface
}
