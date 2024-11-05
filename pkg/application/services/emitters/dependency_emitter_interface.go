package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap ast_maps.AstMap, dependencyList *dependencies.DependencyList)
}
