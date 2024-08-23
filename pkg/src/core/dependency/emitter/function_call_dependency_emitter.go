package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface/TokenReferenceWithDependenciesInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
)

type FunctionCallDependencyEmitter struct{}

func NewFunctionCallDependencyEmitter() *FunctionCallDependencyEmitter {
	return &FunctionCallDependencyEmitter{}
}

func (f *FunctionCallDependencyEmitter) GetName() string {
	return "FunctionCallDependencyEmitter"
}
func (f *FunctionCallDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency.DependencyList) {
	references := make([]TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetClassLikeReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFunctionReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFileReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)
}

func (f *FunctionCallDependencyEmitter) createDependenciesForReferences(references []TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface, astMap ast_map.AstMap, dependencyList *dependency.DependencyList) {
	for _, referenceInterface := range references {
		reference := referenceInterface.(TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface)
		for _, dependencyToken := range reference.GetDependencies() {
			if dependencyToken.Context.DependencyType != DependencyType.DependencyTypeUnresolvedFunctionCall {
				continue
			}
			token := dependencyToken.Token
			dependencyList.AddDependency(dependency.NewDependency(reference.GetToken(), token, dependencyToken.Context))
			functionToken := token.(*ast_map.FunctionToken)
			if functionReference := astMap.GetFunctionReferenceForToken(functionToken); functionReference != nil {
				dependencyList.AddDependency(dependency.NewDependency(reference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
