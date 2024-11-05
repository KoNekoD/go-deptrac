package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassSuperglobalDependencyEmitter struct{}

func NewClassSuperglobalDependencyEmitter() *ClassSuperglobalDependencyEmitter {
	return &ClassSuperglobalDependencyEmitter{}
}

func (c ClassSuperglobalDependencyEmitter) GetName() string {
	return "ClassSuperglobalDependencyEmitter"
}

func (c ClassSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_maps.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType != enums.DependencyTypeSuperGlobalVariable {
				continue
			}
			dependencyList.AddDependency(dependencies.NewDependency(classReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
