package tokens

import (
	tokens3 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type TokenReferenceWithDependenciesInterface interface {
	tokens2.TokenReferenceInterface
	GetDependencies() []*tokens3.DependencyToken
}
