package flatteners

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type InheritanceFlattener struct{}

func NewInheritanceFlattener() *InheritanceFlattener {
	return &InheritanceFlattener{}
}

func (f *InheritanceFlattener) FlattenDependencies(astMap ast_map.AstMap, dependencyList *dependencies.DependencyList) {
	for _, classLikeReference := range astMap.GetClassLikeReferences() {
		classLikeName := classLikeReference.GetToken().(*tokens.ClassLikeToken)
		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			for _, dep := range dependencyList.GetDependenciesByClass(inherit.ClassLikeName.ToString()) {
				dependencyList.AddInheritDependency(dependencies.NewInheritDependency(classLikeName, dep.GetDependent(), dep, inherit))
			}
		}
	}
}
