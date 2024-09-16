package dependency_core

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

type Dependency struct {
	depender  ast_contract2.TokenInterface
	dependent ast_contract2.TokenInterface
	context   *ast_contract2.DependencyContext
}

func NewDependency(depender ast_contract2.TokenInterface, dependent ast_contract2.TokenInterface, context *ast_contract2.DependencyContext) *Dependency {
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

func (d *Dependency) GetDepender() ast_contract2.TokenInterface {
	return d.depender
}

func (d *Dependency) GetDependent() ast_contract2.TokenInterface {
	return d.dependent
}

func (d *Dependency) GetContext() *ast_contract2.DependencyContext {
	return d.context
}
