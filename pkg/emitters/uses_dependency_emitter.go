package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/nodes"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"strings"
)

type UsesDependencyEmitter struct{}

func NewUsesDependencyEmitter() *UsesDependencyEmitter {
	return &UsesDependencyEmitter{}
}

func (u *UsesDependencyEmitter) GetName() string {
	return "UsesDependencyEmitter"
}

func (u *UsesDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	references := make([]tokens.TaggedTokenReferenceInterface, 0)
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

	FQDNIndex := nodes.FQDNIndexNode{}

	for _, reference := range referencesFQDN {
		pathSplit := strings.Split(reference, "\\")
		FQDNIndex.SetNestedNode(pathSplit)
	}

	for _, fileReference := range astMap.GetFileReferences() {
		for _, astStructReference := range fileReference.ClassLikeReferences {
			for _, emittedDependency := range fileReference.Dependencies {
				if emittedDependency.Context.DependencyType == enums.DependencyTypeUse && u.IsFQDN(emittedDependency, FQDNIndex) {
					dependencyList.AddDependency(dependencies.NewDependency(astStructReference.GetToken(), emittedDependency.Token, emittedDependency.Context))
				}
			}
		}
	}

}
func (u *UsesDependencyEmitter) IsFQDN(dependency *tokens.DependencyToken, FQDNIndex nodes.FQDNIndexNode) bool {
	dependencyFQDN := dependency.Token.ToString()
	pathSplit := strings.Split(dependencyFQDN, "\\")
	value := FQDNIndex.GetNestedNode(pathSplit)
	if value == nil {
		return true
	}
	return value.IsFQDN()
}
