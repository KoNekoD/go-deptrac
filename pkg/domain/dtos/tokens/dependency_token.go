package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

type DependencyToken struct {
	Token   TokenInterface
	Context *dependencies.DependencyContext
}

func NewDependencyToken(token TokenInterface, context *dependencies.DependencyContext) *DependencyToken {
	return &DependencyToken{Token: token, Context: context}
}
