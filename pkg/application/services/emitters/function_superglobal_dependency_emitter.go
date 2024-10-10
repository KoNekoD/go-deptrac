package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	dependencies2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FunctionSuperglobalDependencyEmitter struct{}

func NewFunctionSuperglobalDependencyEmitter() *FunctionSuperglobalDependencyEmitter {
	return &FunctionSuperglobalDependencyEmitter{}
}

func (f *FunctionSuperglobalDependencyEmitter) GetName() string {
	return "FunctionSuperglobalDependencyEmitter"
}

func (f *FunctionSuperglobalDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies2.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType != enums.DependencyTypeSuperGlobalVariable {
					continue
				}

				dependencyList.AddDependency(dependencies2.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}