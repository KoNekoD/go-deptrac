package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type ClassSuperglobalDependencyEmitter struct{}

func NewClassSuperglobalDependencyEmitter() *ClassSuperglobalDependencyEmitter {
	return &ClassSuperglobalDependencyEmitter{}
}

func (c ClassSuperglobalDependencyEmitter) GetName() string {
	return "ClassSuperglobalDependencyEmitter"
}

func (c ClassSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType != ast_contract.DependencyTypeSuperGlobalVariable {
				continue
			}
			dependencyList.AddDependency(dependency_core2.NewDependency(classReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
