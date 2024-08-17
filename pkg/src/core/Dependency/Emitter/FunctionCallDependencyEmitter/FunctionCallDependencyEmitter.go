package FunctionCallDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface/TokenReferenceWithDependenciesInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type FunctionCallDependencyEmitter struct{}

func NewFunctionCallDependencyEmitter() *FunctionCallDependencyEmitter {
	return &FunctionCallDependencyEmitter{}
}

func (f *FunctionCallDependencyEmitter) GetName() string {
	return "FunctionCallDependencyEmitter"
}
func (f *FunctionCallDependencyEmitter) ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
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

func (f *FunctionCallDependencyEmitter) createDependenciesForReferences(references []TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface, astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, referenceInterface := range references {
		reference := referenceInterface.(TokenReferenceWithDependenciesInterface.TokenReferenceWithDependenciesInterface)
		for _, dependency := range reference.GetDependencies() {
			if dependency.Context.DependencyType != DependencyType.DependencyTypeUnresolvedFunctionCall {
				continue
			}
			token := dependency.Token
			dependencyList.AddDependency(Dependency.NewDependency(reference.GetToken(), token, dependency.Context))
			functionToken := token.(*AstMap.FunctionToken)
			if functionReference := astMap.GetFunctionReferenceForToken(functionToken); functionReference != nil {
				dependencyList.AddDependency(Dependency.NewDependency(reference.GetToken(), dependency.Token, dependency.Context))
			}
		}
	}
}
