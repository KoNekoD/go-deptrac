package dependencies

import "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"

type DependencyToken struct {
	Token   tokens.TokenInterface
	Context *DependencyContext
}

func NewDependencyToken(token tokens.TokenInterface, context *DependencyContext) *DependencyToken {
	return &DependencyToken{Token: token, Context: context}
}
