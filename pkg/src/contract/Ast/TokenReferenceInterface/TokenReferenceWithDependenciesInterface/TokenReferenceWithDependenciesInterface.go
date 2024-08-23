package TokenReferenceWithDependenciesInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type TokenReferenceWithDependenciesInterface interface {
	TokenReferenceInterface.TokenReferenceInterface
	GetDependencies() []*ast_map.DependencyToken
}
