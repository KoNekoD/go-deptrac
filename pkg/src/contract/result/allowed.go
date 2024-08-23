package result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/dependency"
)

// Allowed - Represents a dependency that is allowed to exist given the defined rules
type Allowed struct {
	Dependency     dependency.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewAllowed(dependency dependency.DependencyInterface, dependerLayer string, dependentLayer string) *Allowed {
	return &Allowed{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (a *Allowed) GetDependency() dependency.DependencyInterface {
	return a.Dependency
}

func (a *Allowed) GetDependerLayer() string {
	return a.DependerLayer
}

func (a *Allowed) GetDependentLayer() string {
	return a.DependentLayer
}
