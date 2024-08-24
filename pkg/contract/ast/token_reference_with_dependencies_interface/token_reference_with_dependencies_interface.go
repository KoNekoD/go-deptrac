package token_reference_with_dependencies_interface

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
)

type TokenReferenceWithDependenciesInterface interface {
	ast.TokenReferenceInterface
	GetDependencies() []*ast_map.DependencyToken
}
