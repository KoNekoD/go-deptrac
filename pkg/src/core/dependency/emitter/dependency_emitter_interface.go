package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency.DependencyList)
}
