package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FunctionDependencyEmitter struct{}

func NewFunctionDependencyEmitter() *FunctionDependencyEmitter {
	return &FunctionDependencyEmitter{}
}

func (f FunctionDependencyEmitter) GetName() string {
	return "FunctionDependencyEmitter"
}

func (f FunctionDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, astFunctionReference := range fileReference.FunctionReferences {
			for _, dependencyToken := range astFunctionReference.Dependencies {
				if dependencyToken.Context.DependencyType == enums.DependencyTypeSuperGlobalVariable {
					continue
				}

				if dependencyToken.Context.DependencyType == enums.DependencyTypeUnresolvedFunctionCall {
					continue
				}

				dependencyList.AddDependency(dependencies.NewDependency(astFunctionReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
