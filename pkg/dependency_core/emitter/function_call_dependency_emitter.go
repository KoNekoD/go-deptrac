package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract/token_reference_with_dependencies_interface"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type FunctionCallDependencyEmitter struct{}

func NewFunctionCallDependencyEmitter() *FunctionCallDependencyEmitter {
	return &FunctionCallDependencyEmitter{}
}

func (f *FunctionCallDependencyEmitter) GetName() string {
	return "FunctionCallDependencyEmitter"
}
func (f *FunctionCallDependencyEmitter) ApplyDependencies(astMap ast_map2.AstMap, dependencyList *dependency_core2.DependencyList) {
	references := make([]token_reference_with_dependencies_interface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetClassLikeReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]token_reference_with_dependencies_interface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFunctionReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)

	references = make([]token_reference_with_dependencies_interface.TokenReferenceWithDependenciesInterface, 0)
	for _, reference := range astMap.GetFileReferences() {
		references = append(references, reference)
	}
	f.createDependenciesForReferences(references, astMap, dependencyList)
}

func (f *FunctionCallDependencyEmitter) createDependenciesForReferences(references []token_reference_with_dependencies_interface.TokenReferenceWithDependenciesInterface, astMap ast_map2.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, referenceInterface := range references {
		reference := referenceInterface.(token_reference_with_dependencies_interface.TokenReferenceWithDependenciesInterface)
		for _, dependencyToken := range reference.GetDependencies() {
			if dependencyToken.Context.DependencyType != ast_contract.DependencyTypeUnresolvedFunctionCall {
				continue
			}
			token := dependencyToken.Token
			dependencyList.AddDependency(dependency_core2.NewDependency(reference.GetToken(), token, dependencyToken.Context))
			functionToken := token.(*ast_map2.FunctionToken)
			if functionReference := astMap.GetFunctionReferenceForToken(functionToken); functionReference != nil {
				dependencyList.AddDependency(dependency_core2.NewDependency(reference.GetToken(), dependencyToken.Token, dependencyToken.Context))
			}
		}
	}
}
