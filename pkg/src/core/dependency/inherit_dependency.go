package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type InheritDependency struct {
	depender           *ast_map.ClassLikeToken
	dependent          TokenInterface.TokenInterface
	originalDependency DependencyInterface.DependencyInterface
	inheritPath        *ast_map.AstInherit
}

func NewInheritDependency(depender *ast_map.ClassLikeToken, dependent TokenInterface.TokenInterface, originalDependency DependencyInterface.DependencyInterface, inheritPath *ast_map.AstInherit) *InheritDependency {
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

func (i *InheritDependency) GetDepender() TokenInterface.TokenInterface {
	return i.depender
}

func (i *InheritDependency) GetDependent() TokenInterface.TokenInterface {
	return i.dependent
}

func (i *InheritDependency) GetContext() *DependencyContext.DependencyContext {
	return i.originalDependency.GetContext()
}
