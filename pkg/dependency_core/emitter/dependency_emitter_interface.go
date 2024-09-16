package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core.DependencyList)
}
