package result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/dependency"
)

// Uncovered - Represents a dependency that is NOT covered by the current configuration.
type Uncovered struct {
	Dependency dependency.DependencyInterface
	Layer      string
}

func NewUncovered(dependency dependency.DependencyInterface, layer string) *Uncovered {
	return &Uncovered{
		Dependency: dependency,
		Layer:      layer,
	}
}

func (u *Uncovered) GetDependency() dependency.DependencyInterface {
	return u.Dependency
}
