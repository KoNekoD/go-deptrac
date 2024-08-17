package InheritanceFlattener

import (
	AstMap2 "github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InheritDependency"
)

type InheritanceFlattener struct{}

func NewInheritanceFlattener() *InheritanceFlattener {
	return &InheritanceFlattener{}
}

func (f *InheritanceFlattener) FlattenDependencies(astMap AstMap2.AstMap, dependencyList *DependencyList.DependencyList) {
	for _, classLikeReference := range astMap.GetClassLikeReferences() {
		classLikeName := classLikeReference.GetToken().(*AstMap2.ClassLikeToken)
		for _, inherit := range astMap.GetClassInherits(classLikeName) {
			for _, dep := range dependencyList.GetDependenciesByClass(inherit.ClassLikeName.ToString()) {
				dependencyList.AddInheritDependency(InheritDependency.NewInheritDependency(classLikeName, dep.GetDependent(), dep, inherit))
			}
		}
	}
}
