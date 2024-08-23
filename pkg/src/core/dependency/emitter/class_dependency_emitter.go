package emitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	AstMap2 "github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
)

type ClassDependencyEmitter struct{}

func NewClassDependencyEmitter() *ClassDependencyEmitter {
	return &ClassDependencyEmitter{}
}

func (c *ClassDependencyEmitter) GetName() string {
	return "ClassDependencyEmitter"
}

func (c *ClassDependencyEmitter) ApplyDependencies(astMap AstMap2.AstMap, dependencyList *dependency.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*AstMap2.ClassLikeToken)

		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType == DependencyType.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependencyToken.Context.DependencyType == DependencyType.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependency.NewDependency(classLikeName, dependencyToken.Token, dependencyToken.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(dependency.NewDependency(classLikeName, inherit.ClassLikeName, DependencyContext.NewDependencyContext(inherit.FileOccurrence, DependencyType.DependencyTypeInherit)))
		}
	}
}
