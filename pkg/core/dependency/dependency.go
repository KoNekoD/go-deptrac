package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
)

type Dependency struct {
	depender  ast.TokenInterface
	dependent ast.TokenInterface
	context   *ast.DependencyContext
}

func NewDependency(depender ast.TokenInterface, dependent ast.TokenInterface, context *ast.DependencyContext) *Dependency {
	if dependent.ToString() == "" {
		panic("1")
	}

	return &Dependency{depender: depender, dependent: dependent, context: context}
}

func (d *Dependency) Serialize() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name": d.dependent.ToString(),
			"line": d.context.FileOccurrence.Line,
		},
	}
}

func (d *Dependency) GetDepender() ast.TokenInterface {
	return d.depender
}

func (d *Dependency) GetDependent() ast.TokenInterface {
	return d.dependent
}

func (d *Dependency) GetContext() *ast.DependencyContext {
	return d.context
}
