package dependency_core

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type InheritanceFlattener struct{}

func NewInheritanceFlattener() *InheritanceFlattener {
	return &InheritanceFlattener{}
}

func (f *InheritanceFlattener) FlattenDependencies(astMap ast_map.AstMap, dependencyList *DependencyList) {
	for _, classLikeReference := range astMap.GetClassLikeReferences() {
		classLikeName := classLikeReference.GetToken().(*ast_map.ClassLikeToken)
		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			for _, dep := range dependencyList.GetDependenciesByClass(inherit.ClassLikeName.ToString()) {
				dependencyList.AddInheritDependency(NewInheritDependency(classLikeName, dep.GetDependent(), dep, inherit))
			}
		}
	}
}
