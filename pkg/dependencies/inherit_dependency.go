package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type InheritDependency struct {
	depender           *tokens.ClassLikeToken
	dependent          tokens.TokenInterface
	originalDependency DependencyInterface
	inheritPath        *ast_map.AstInherit
}

func NewInheritDependency(depender *tokens.ClassLikeToken, dependent tokens.TokenInterface, originalDependency DependencyInterface, inheritPath *ast_map.AstInherit) *InheritDependency {
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

func (i *InheritDependency) GetDepender() tokens.TokenInterface {
	return i.depender
}

func (i *InheritDependency) GetDependent() tokens.TokenInterface {
	return i.dependent
}

func (i *InheritDependency) GetContext() *DependencyContext {
	return i.originalDependency.GetContext()
}
