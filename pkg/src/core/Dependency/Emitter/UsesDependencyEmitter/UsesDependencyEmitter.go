package UsesDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TaggedTokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/FQDNIndexNode"
	"strings"
)

type UsesDependencyEmitter struct{}

func NewUsesDependencyEmitter() *UsesDependencyEmitter {
	return &UsesDependencyEmitter{}
}

func (u *UsesDependencyEmitter) GetName() string {
	return "UsesDependencyEmitter"
}

func (u *UsesDependencyEmitter) ApplyDependencies(astMap AstMap.AstMap, dependencyList *DependencyList.DependencyList) {
	references := make([]TaggedTokenReferenceInterface.TaggedTokenReferenceInterface, 0)
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

	FQDNIndex := FQDNIndexNode.FQDNIndexNode{}

	for _, reference := range referencesFQDN {
		pathSplit := strings.Split(reference, "\\")
		FQDNIndex.SetNestedNode(pathSplit)
	}

	for _, fileReference := range astMap.GetFileReferences() {
		for _, astStructReference := range fileReference.ClassLikeReferences {
			for _, emittedDependency := range fileReference.Dependencies {
				if emittedDependency.Context.DependencyType == DependencyType.DependencyTypeUse && u.IsFQDN(emittedDependency, FQDNIndex) {
					dependencyList.AddDependency(Dependency.NewDependency(astStructReference.GetToken(), emittedDependency.Token, emittedDependency.Context))
				}
			}
		}
	}

}
func (u *UsesDependencyEmitter) IsFQDN(dependency *AstMap.DependencyToken, FQDNIndex FQDNIndexNode.FQDNIndexNode) bool {
	dependencyFQDN := dependency.Token.ToString()
	pathSplit := strings.Split(dependencyFQDN, "\\")
	value := FQDNIndex.GetNestedNode(pathSplit)
	if value == nil {
		return true
	}
	return value.IsFQDN()
}
