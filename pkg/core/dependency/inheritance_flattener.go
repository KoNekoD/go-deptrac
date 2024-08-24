package dependency

import (
	AstMap2 "github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
)

type InheritanceFlattener struct{}

func NewInheritanceFlattener() *InheritanceFlattener {
	return &InheritanceFlattener{}
}

func (f *InheritanceFlattener) FlattenDependencies(astMap AstMap2.AstMap, dependencyList *DependencyList) {
	for _, classLikeReference := range astMap.GetClassLikeReferences() {
		classLikeName := classLikeReference.GetToken().(*AstMap2.ClassLikeToken)
		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			for _, dep := range dependencyList.GetDependenciesByClass(inherit.ClassLikeName.ToString()) {
				dependencyList.AddInheritDependency(NewInheritDependency(classLikeName, dep.GetDependent(), dep, inherit))
			}
		}
	}
}
