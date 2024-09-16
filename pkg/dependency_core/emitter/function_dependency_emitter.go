package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type FunctionDependencyEmitter struct{}

func NewFunctionDependencyEmitter() *FunctionDependencyEmitter {
	return &FunctionDependencyEmitter{}
}

func (f FunctionDependencyEmitter) GetName() string {
	return "FunctionDependencyEmitter"
}

func (f FunctionDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType == ast_contract.DependencyTypeSuperGlobalVariable {
					continue
				}

				if dependencyToken.Context.DependencyType == ast_contract.DependencyTypeUnresolvedFunctionCall {
					continue
				}

				dependencyList.AddDependency(dependency_core2.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
