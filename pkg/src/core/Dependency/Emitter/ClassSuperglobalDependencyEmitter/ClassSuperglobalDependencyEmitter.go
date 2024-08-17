package ClassSuperglobalDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type ClassSuperglobalDependencyEmitter struct{}

func NewClassSuperglobalDependencyEmitter() *ClassSuperglobalDependencyEmitter {
	return &ClassSuperglobalDependencyEmitter{}
}

func (c ClassSuperglobalDependencyEmitter) GetName() string {
	return "ClassSuperglobalDependencyEmitter"
}

func (c ClassSuperglobalDependencyEmitter) ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		for _, dependency := range classReference.Dependencies {
			if dependency.Context.DependencyType != DependencyType.DependencyTypeSuperGlobalVariable {
				continue
			}
			dependencyList.AddDependency(Dependency.NewDependency(classReference.GetToken(), dependency.Token, dependency.Context))
		}
	}
}
