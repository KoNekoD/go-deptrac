package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency"
)

type ClassSuperglobalDependencyEmitter struct{}

func NewClassSuperglobalDependencyEmitter() *ClassSuperglobalDependencyEmitter {
	return &ClassSuperglobalDependencyEmitter{}
}

func (c ClassSuperglobalDependencyEmitter) GetName() string {
	return "ClassSuperglobalDependencyEmitter"
}

func (c ClassSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType != ast.DependencyTypeSuperGlobalVariable {
				continue
			}
			dependencyList.AddDependency(dependency.NewDependency(classReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
