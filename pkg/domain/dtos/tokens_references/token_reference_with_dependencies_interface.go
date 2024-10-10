package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

type TokenReferenceWithDependenciesInterface interface {
	TokenReferenceInterface
	GetDependencies() []*tokens.DependencyToken
}
