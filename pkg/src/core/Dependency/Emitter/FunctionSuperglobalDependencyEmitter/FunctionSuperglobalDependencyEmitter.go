package FunctionSuperglobalDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type FunctionSuperglobalDependencyEmitter struct{}

func NewFunctionSuperglobalDependencyEmitter() *FunctionSuperglobalDependencyEmitter {
	return &FunctionSuperglobalDependencyEmitter{}
}

func (f *FunctionSuperglobalDependencyEmitter) GetName() string {
	return "FunctionSuperglobalDependencyEmitter"
}

func (f *FunctionSuperglobalDependencyEmitter) ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependency := range astFunctionReference.Dependencies {
				if dependency.Context.DependencyType != DependencyType.DependencyTypeSuperGlobalVariable {
					continue
				}

				dependencyList.AddDependency(Dependency.NewDependency(astFunctionReference.GetToken(), dependency.Token, dependency.Context))
			}
		}
	}
}
