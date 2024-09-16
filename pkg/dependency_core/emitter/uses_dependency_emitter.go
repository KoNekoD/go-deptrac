package emitter

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"strings"
)

type UsesDependencyEmitter struct{}

func NewUsesDependencyEmitter() *UsesDependencyEmitter {
	return &UsesDependencyEmitter{}
}

func (u *UsesDependencyEmitter) GetName() string {
	return "UsesDependencyEmitter"
}

func (u *UsesDependencyEmitter) ApplyDependencies(astMap ast_map2.AstMap, dependencyList *dependency_core2.DependencyList) {
	references := make([]ast_contract2.TaggedTokenReferenceInterface, 0)
	for _, structLikeReference := range astMap.GetClassLikeReferences() {
		references = append(references, structLikeReference)
	}
	for _, functionReference := range astMap.GetFunctionReferences() {
		references = append(references, functionReference)
	}

	referencesFQDN := make([]string, 0)
	for _, reference := range references {
		referencesFQDN = append(referencesFQDN, reference.GetToken().ToString())
	}

	FQDNIndex := FQDNIndexNode{}

	for _, reference := range referencesFQDN {
		pathSplit := strings.Split(reference, "\\")
		FQDNIndex.SetNestedNode(pathSplit)
	}

	for _, fileReference := range astMap.GetFileReferences() {
		for _, astStructReference := range fileReference.ClassLikeReferences {
			for _, emittedDependency := range fileReference.Dependencies {
				if emittedDependency.Context.DependencyType == ast_contract2.DependencyTypeUse && u.IsFQDN(emittedDependency, FQDNIndex) {
					dependencyList.AddDependency(dependency_core2.NewDependency(astStructReference.GetToken(), emittedDependency.Token, emittedDependency.Context))
				}
			}
		}
	}

}
func (u *UsesDependencyEmitter) IsFQDN(dependency *ast_map2.DependencyToken, FQDNIndex FQDNIndexNode) bool {
	dependencyFQDN := dependency.Token.ToString()
	pathSplit := strings.Split(dependencyFQDN, "\\")
	value := FQDNIndex.GetNestedNode(pathSplit)
	if value == nil {
		return true
	}
	return value.IsFQDN()
}
