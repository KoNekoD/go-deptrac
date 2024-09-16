package ast_contract

// TaggedTokenReferenceInterface - Represents the AST-TokenInterface, its location, and associated tags.
type TaggedTokenReferenceInterface interface {
	TokenReferenceInterface
	HasTag(name string) bool
	GetTagLines(name string) []string
}
