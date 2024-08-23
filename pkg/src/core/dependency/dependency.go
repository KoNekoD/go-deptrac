package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
)

type Dependency struct {
	depender  TokenInterface.TokenInterface
	dependent TokenInterface.TokenInterface
	context   *DependencyContext.DependencyContext
}

func NewDependency(depender TokenInterface.TokenInterface, dependent TokenInterface.TokenInterface, context *DependencyContext.DependencyContext) *Dependency {
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

func (d *Dependency) GetDepender() TokenInterface.TokenInterface {
	return d.depender
}

func (d *Dependency) GetDependent() TokenInterface.TokenInterface {
	return d.dependent
}

func (d *Dependency) GetContext() *DependencyContext.DependencyContext {
	return d.context
}
