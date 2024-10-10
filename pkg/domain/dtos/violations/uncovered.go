package violations

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

// Uncovered - Represents a dependency_contract that is NOT covered by the current configuration.
type Uncovered struct {
	Dependency dependencies.DependencyInterface
	Layer      string
}

func NewUncovered(dependency dependencies.DependencyInterface, layer string) *Uncovered {
	return &Uncovered{
		Dependency: dependency,
		Layer:      layer,
	}
}

func (u *Uncovered) GetDependency() dependencies.DependencyInterface {
	return u.Dependency
}
