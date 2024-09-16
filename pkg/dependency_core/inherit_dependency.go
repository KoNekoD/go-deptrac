package dependency_core

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

type InheritDependency struct {
	depender           *ast_map2.ClassLikeToken
	dependent          ast_contract2.TokenInterface
	originalDependency Dependency2.DependencyInterface
	inheritPath        *ast_map2.AstInherit
}

func NewInheritDependency(depender *ast_map2.ClassLikeToken, dependent ast_contract2.TokenInterface, originalDependency Dependency2.DependencyInterface, inheritPath *ast_map2.AstInherit) *InheritDependency {
	return &InheritDependency{depender: depender, dependent: dependent, originalDependency: originalDependency, inheritPath: inheritPath}
}

func (i *InheritDependency) Serialize() []map[string]interface{} {
	var buffer []map[string]interface{}

	var path []map[string]interface{}
	for _, p := range i.inheritPath.GetPath() {
		path = append(path, map[string]interface{}{"name": p.ClassLikeName.ToString(), "line": p.FileOccurrence.Line})
	}

	buffer = append(buffer, path...)

	buffer = append(buffer, map[string]interface{}{"name": i.inheritPath.ClassLikeName.ToString(), "line": i.inheritPath.FileOccurrence.Line})
	buffer = append(buffer, map[string]interface{}{"name": i.originalDependency.GetDependent().ToString(), "line": i.originalDependency.GetContext().FileOccurrence.Line})

	return buffer
}

func (i *InheritDependency) GetDepender() ast_contract2.TokenInterface {
	return i.depender
}

func (i *InheritDependency) GetDependent() ast_contract2.TokenInterface {
	return i.dependent
}

func (i *InheritDependency) GetContext() *ast_contract2.DependencyContext {
	return i.originalDependency.GetContext()
}
