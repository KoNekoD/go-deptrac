package Uncovered

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"

// Uncovered - Represents a dependency that is NOT covered by the current configuration.
type Uncovered struct {
	Dependency DependencyInterface.DependencyInterface
	Layer      string
}

func NewUncovered(dependency DependencyInterface.DependencyInterface, layer string) *Uncovered {
	return &Uncovered{
		Dependency: dependency,
		Layer:      layer,
	}
}

func (u *Uncovered) GetDependency() DependencyInterface.DependencyInterface {
	return u.Dependency
}
