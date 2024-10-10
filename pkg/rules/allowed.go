package rules

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

// Allowed - Represents a dependency_contract that is allowed to exist given the defined rules
type Allowed struct {
	Dependency     dependencies.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewAllowed(dependency dependencies.DependencyInterface, dependerLayer string, dependentLayer string) *Allowed {
	return &Allowed{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (a *Allowed) GetDependency() dependencies.DependencyInterface {
	return a.Dependency
}

func (a *Allowed) GetDependerLayer() string {
	return a.DependerLayer
}

func (a *Allowed) GetDependentLayer() string {
	return a.DependentLayer
}
