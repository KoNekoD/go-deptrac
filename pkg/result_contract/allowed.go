package result_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// Allowed - Represents a dependency_contract that is allowed to exist given the defined rules
type Allowed struct {
	Dependency     dependency_contract.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewAllowed(dependency dependency_contract.DependencyInterface, dependerLayer string, dependentLayer string) *Allowed {
	return &Allowed{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (a *Allowed) GetDependency() dependency_contract.DependencyInterface {
	return a.Dependency
}

func (a *Allowed) GetDependerLayer() string {
	return a.DependerLayer
}

func (a *Allowed) GetDependentLayer() string {
	return a.DependentLayer
}
