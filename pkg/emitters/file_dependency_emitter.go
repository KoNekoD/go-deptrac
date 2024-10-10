package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FileDependencyEmitter struct{}

func NewFileDependencyEmitter() *FileDependencyEmitter {
	return &FileDependencyEmitter{}
}

func (f FileDependencyEmitter) GetName() string {
	return "FileDependencyEmitter"
}

func (f FileDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, dependencyToken := range fileReference.Dependencies {
			if dependencyToken.Context.DependencyType == enums.DependencyTypeUse {
				continue
			}

			if dependencyToken.Context.DependencyType == enums.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependencies.NewDependency(fileReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
