package FileDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type FileDependencyEmitter struct{}

func NewFileDependencyEmitter() *FileDependencyEmitter {
	return &FileDependencyEmitter{}
}

func (f FileDependencyEmitter) GetName() string {
	return "FileDependencyEmitter"
}

func (f FileDependencyEmitter) ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, dependency := range fileReference.Dependencies {
			if dependency.Context.DependencyType == DependencyType.DependencyTypeUse {
				continue
			}

			if dependency.Context.DependencyType == DependencyType.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(Dependency.NewDependency(fileReference.GetToken(), dependency.Token, dependency.Context))
		}
	}
}
