package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type FunctionSuperglobalDependencyEmitter struct{}

func NewFunctionSuperglobalDependencyEmitter() *FunctionSuperglobalDependencyEmitter {
	return &FunctionSuperglobalDependencyEmitter{}
}

func (f *FunctionSuperglobalDependencyEmitter) GetName() string {
	return "FunctionSuperglobalDependencyEmitter"
}

func (f *FunctionSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType != ast_contract.DependencyTypeSuperGlobalVariable {
					continue
				}

				dependencyList.AddDependency(dependency_core2.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
