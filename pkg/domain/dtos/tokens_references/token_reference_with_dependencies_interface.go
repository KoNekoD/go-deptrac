package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

type TokenReferenceWithDependenciesInterface interface {
	TokenReferenceInterface
	GetDependencies() []*dependencies.DependencyToken
}
