package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassSuperglobalDependencyEmitter struct{}

func NewClassSuperglobalDependencyEmitter() *ClassSuperglobalDependencyEmitter {
	return &ClassSuperglobalDependencyEmitter{}
}

func (c ClassSuperglobalDependencyEmitter) GetName() string {
	return "ClassSuperglobalDependencyEmitter"
}

func (c ClassSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType != enums.DependencyTypeSuperGlobalVariable {
				continue
			}
			dependencyList.AddDependency(dependencies.NewDependency(classReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
