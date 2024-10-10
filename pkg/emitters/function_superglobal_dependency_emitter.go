package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
)

type FunctionSuperglobalDependencyEmitter struct{}

func NewFunctionSuperglobalDependencyEmitter() *FunctionSuperglobalDependencyEmitter {
	return &FunctionSuperglobalDependencyEmitter{}
}

func (f *FunctionSuperglobalDependencyEmitter) GetName() string {
	return "FunctionSuperglobalDependencyEmitter"
}

func (f *FunctionSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType != dependencies.DependencyTypeSuperGlobalVariable {
					continue
				}

				dependencyList.AddDependency(dependencies.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
