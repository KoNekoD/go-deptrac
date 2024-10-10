package dependencies

import (
	dependencies2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

type Dependency struct {
	depender  tokens.TokenInterface
	dependent tokens.TokenInterface
	context   *dependencies2.DependencyContext
}

func NewDependency(depender tokens.TokenInterface, dependent tokens.TokenInterface, context *dependencies2.DependencyContext) *Dependency {
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

func (d *Dependency) GetDepender() tokens.TokenInterface {
	return d.depender
}

func (d *Dependency) GetDependent() tokens.TokenInterface {
	return d.dependent
}

func (d *Dependency) GetContext() *dependencies2.DependencyContext {
	return d.context
}
