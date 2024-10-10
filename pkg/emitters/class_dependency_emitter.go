package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type ClassDependencyEmitter struct{}

func NewClassDependencyEmitter() *ClassDependencyEmitter {
	return &ClassDependencyEmitter{}
}

func (c *ClassDependencyEmitter) GetName() string {
	return "ClassDependencyEmitter"
}

func (c *ClassDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*tokens.ClassLikeToken)

		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType == dependencies.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependencyToken.Context.DependencyType == dependencies.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependencies.NewDependency(classLikeName, dependencyToken.Token, dependencyToken.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(dependencies.NewDependency(classLikeName, inherit.ClassLikeName, dependencies.NewDependencyContext(inherit.FileOccurrence, dependencies.DependencyTypeInherit)))
		}
	}
}