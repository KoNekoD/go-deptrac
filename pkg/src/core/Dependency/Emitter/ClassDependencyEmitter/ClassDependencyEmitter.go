package ClassDependencyEmitter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	AstMap2 "github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
)

type ClassDependencyEmitter struct{}

func NewClassDependencyEmitter() *ClassDependencyEmitter {
	return &ClassDependencyEmitter{}
}

func (c *ClassDependencyEmitter) GetName() string {
	return "ClassDependencyEmitter"
}

func (c *ClassDependencyEmitter) ApplyDependencies(astMap AstMap2.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*AstMap2.ClassLikeToken)

		for _, dependency := range classReference.Dependencies {
			if dependency.Context.DependencyType == DependencyType.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependency.Context.DependencyType == DependencyType.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(Dependency.NewDependency(classLikeName, dependency.Token, dependency.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(Dependency.NewDependency(classLikeName, inherit.ClassLikeName, DependencyContext.NewDependencyContext(inherit.FileOccurrence, DependencyType.DependencyTypeInherit)))
		}
	}
}
