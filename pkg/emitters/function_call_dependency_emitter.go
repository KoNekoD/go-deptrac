package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type FunctionCallDependencyEmitter struct{}

func NewFunctionCallDependencyEmitter() *FunctionCallDependencyEmitter {
	return &FunctionCallDependencyEmitter{}
}

func (f *FunctionCallDependencyEmitter) GetName() string {
	return "FunctionCallDependencyEmitter"
}
func (f *FunctionCallDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	references := make([]tokens.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetClassLikeReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]tokens.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFunctionReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]tokens.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFileReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)
}

func (f *FunctionCallDependencyEmitter) createDependenciesForReferences(references []tokens.TokenReferenceWithDependenciesInterface, astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, referenceInterface := range references {
		reference := referenceInterface.(tokens.TokenReferenceWithDependenciesInterface)
		for _, dependencyToken := range reference.GetDependencies() {
			if dependencyToken.Context.DependencyType != dependencies.DependencyTypeUnresolvedFunctionCall {
				continue
			}
			token := dependencyToken.Token
			dependencyList.AddDependency(dependencies.NewDependency(reference.GetToken(), token, dependencyToken.Context))
			functionToken := token.(*tokens.FunctionToken)
			if functionReference := astMap.GetFunctionReferenceForToken(functionToken); functionReference != nil {
				dependencyList.AddDependency(dependencies.NewDependency(reference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
