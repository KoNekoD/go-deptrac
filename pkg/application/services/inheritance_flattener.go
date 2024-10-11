package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

type InheritanceFlattener struct{}

func NewInheritanceFlattener() *InheritanceFlattener {
	return &InheritanceFlattener{}
}

func (f *InheritanceFlattener) FlattenDependencies(astMap ast_maps.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classLikeReference := range astMap.GetClassLikeReferences() {
		classLikeName := classLikeReference.GetToken().(*tokens.ClassLikeToken)
		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			for _, dep := range dependencyList.GetDependenciesByClass(inherit.ClassLikeName.ToString()) {
				dependencyList.AddInheritDependency(dependencies.NewInheritDependency(classLikeName, dep.GetDependent(), dep, inherit))
			}
		}
	}
}
