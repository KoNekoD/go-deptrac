package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList)
}
