package result_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// Uncovered - Represents a dependency_contract that is NOT covered by the current configuration.
type Uncovered struct {
	Dependency dependency_contract.DependencyInterface
	Layer      string
}

func NewUncovered(dependency dependency_contract.DependencyInterface, layer string) *Uncovered {
	return &Uncovered{
		Dependency: dependency,
		Layer:      layer,
	}
}

func (u *Uncovered) GetDependency() dependency_contract.DependencyInterface {
	return u.Dependency
}
