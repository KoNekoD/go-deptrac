package token_reference_with_dependencies_interface

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type TokenReferenceWithDependenciesInterface interface {
	ast_contract.TokenReferenceInterface
	GetDependencies() []*ast_map.DependencyToken
}
