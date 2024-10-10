package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	dependencies2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassDependencyEmitter struct{}

func NewClassDependencyEmitter() *ClassDependencyEmitter {
	return &ClassDependencyEmitter{}
}

func (c *ClassDependencyEmitter) GetName() string {
	return "ClassDependencyEmitter"
}

func (c *ClassDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependencies2.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*tokens.ClassLikeToken)

		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType == enums.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependencyToken.Context.DependencyType == enums.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependencies2.NewDependency(classLikeName, dependencyToken.Token, dependencyToken.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(dependencies2.NewDependency(classLikeName, inherit.ClassLikeName, dependencies2.NewDependencyContext(inherit.FileOccurrence, enums.DependencyTypeInherit)))
		}
	}
}
