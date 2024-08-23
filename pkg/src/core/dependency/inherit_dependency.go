package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type InheritDependency struct {
	depender           *ast_map.ClassLikeToken
	dependent          ast.TokenInterface
	originalDependency Dependency2.DependencyInterface
	inheritPath        *ast_map.AstInherit
}

func NewInheritDependency(depender *ast_map.ClassLikeToken, dependent ast.TokenInterface, originalDependency Dependency2.DependencyInterface, inheritPath *ast_map.AstInherit) *InheritDependency {
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

func (i *InheritDependency) GetDepender() ast.TokenInterface {
	return i.depender
}

func (i *InheritDependency) GetDependent() ast.TokenInterface {
	return i.dependent
}

func (i *InheritDependency) GetContext() *ast.DependencyContext {
	return i.originalDependency.GetContext()
}
