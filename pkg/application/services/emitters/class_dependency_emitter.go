package emitters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
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

func (c *ClassDependencyEmitter) ApplyDependencies(astMap ast_maps.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*tokens.ClassLikeToken)

		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType == enums.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependencyToken.Context.DependencyType == enums.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependencies.NewDependency(classLikeName, dependencyToken.Token, dependencyToken.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(dependencies.NewDependency(classLikeName, inherit.ClassLikeName, dependencies.NewDependencyContext(inherit.FileOccurrence, enums.DependencyTypeInherit)))
		}
	}
}
