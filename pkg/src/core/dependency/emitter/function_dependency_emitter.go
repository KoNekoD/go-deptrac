package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
)

type FunctionDependencyEmitter struct{}

func NewFunctionDependencyEmitter() *FunctionDependencyEmitter {
	return &FunctionDependencyEmitter{}
}

func (f FunctionDependencyEmitter) GetName() string {
	return "FunctionDependencyEmitter"
}

func (f FunctionDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType == ast.DependencyTypeSuperGlobalVariable {
					continue
				}

				if dependencyToken.Context.DependencyType == ast.DependencyTypeUnresolvedFunctionCall {
					continue
				}

				dependencyList.AddDependency(dependency.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
