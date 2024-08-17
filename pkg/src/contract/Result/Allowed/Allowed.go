package Allowed

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"

// Allowed - Represents a dependency that is allowed to exist given the defined rules
type Allowed struct {
	Dependency     DependencyInterface.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewAllowed(dependency DependencyInterface.DependencyInterface, dependerLayer string, dependentLayer string) *Allowed {
	return &Allowed{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (a *Allowed) GetDependency() DependencyInterface.DependencyInterface {
	return a.Dependency
}

func (a *Allowed) GetDependerLayer() string {
	return a.DependerLayer
}

func (a *Allowed) GetDependentLayer() string {
	return a.DependentLayer
}
