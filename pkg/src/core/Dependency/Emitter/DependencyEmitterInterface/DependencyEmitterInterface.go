package DependencyEmitterInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type DependencyEmitterInterface interface {
	GetName() string
	ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList)
}
