package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList)
}
