package TaggedTokenReferenceInterface

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"

// TaggedTokenReferenceInterface - Represents the AST-TokenInterface, its location, and associated tags.
type TaggedTokenReferenceInterface interface {
	TokenReferenceInterface.TokenReferenceInterface
	HasTag(name string) bool
	GetTagLines(name string) []string
}
