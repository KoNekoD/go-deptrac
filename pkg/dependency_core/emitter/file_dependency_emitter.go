package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type FileDependencyEmitter struct{}

func NewFileDependencyEmitter() *FileDependencyEmitter {
	return &FileDependencyEmitter{}
}

func (f FileDependencyEmitter) GetName() string {
	return "FileDependencyEmitter"
}

func (f FileDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, fileReference := range astMap.GetFileReferences() {
		for _, dependencyToken := range fileReference.Dependencies {
			if dependencyToken.Context.DependencyType == ast_contract.DependencyTypeUse {
				continue
			}

			if dependencyToken.Context.DependencyType == ast_contract.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependency_core2.NewDependency(fileReference.GetToken(), dependencyToken.Token, dependencyToken.Context))
		}
	}
}
