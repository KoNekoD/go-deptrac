package token_reference_with_dependencies_interface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type TokenReferenceWithDependenciesInterface interface {
	ast.TokenReferenceInterface
	GetDependencies() []*ast_map.DependencyToken
}
