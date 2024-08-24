package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency.DependencyList)
}
