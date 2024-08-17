package TokenReferenceWithDependenciesInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
)

type TokenReferenceWithDependenciesInterface interface {
	TokenReferenceInterface.TokenReferenceInterface
	GetDependencies() []*AstMap.DependencyToken
}
