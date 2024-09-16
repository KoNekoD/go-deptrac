package emitter

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
)

type ClassDependencyEmitter struct{}

func NewClassDependencyEmitter() *ClassDependencyEmitter {
	return &ClassDependencyEmitter{}
}

func (c *ClassDependencyEmitter) GetName() string {
	return "ClassDependencyEmitter"
}

func (c *ClassDependencyEmitter) ApplyDependencies(astMap ast_map.AstMap, dependencyList *dependency_core2.DependencyList) {
	for _, classReference := range astMap.GetClassLikeReferences() {
		classLikeName := classReference.GetToken().(*ast_map.ClassLikeToken)

		for _, dependencyToken := range classReference.Dependencies {
			if dependencyToken.Context.DependencyType == ast_contract2.DependencyTypeSuperGlobalVariable {
				continue
			}

			if dependencyToken.Context.DependencyType == ast_contract2.DependencyTypeUnresolvedFunctionCall {
				continue
			}

			dependencyList.AddDependency(dependency_core2.NewDependency(classLikeName, dependencyToken.Token, dependencyToken.Context))
		}

		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			dependencyList.AddDependency(dependency_core2.NewDependency(classLikeName, inherit.ClassLikeName, ast_contract2.NewDependencyContext(inherit.FileOccurrence, ast_contract2.DependencyTypeInherit)))
		}
	}
}
