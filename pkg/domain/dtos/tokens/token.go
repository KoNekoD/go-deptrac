package tokens

// TokenInterface - Represents an AST-TokenInterface, which can be referenced as dependency_contract.
type TokenInterface interface {
	ToString() string
	tokenInterface()
}
