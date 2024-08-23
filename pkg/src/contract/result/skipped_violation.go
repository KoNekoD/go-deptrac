package result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/dependency"
)

// SkippedViolation - Represents a Violation that is being skipped by the baseline file
type SkippedViolation struct {
	Dependency     dependency.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewSkippedViolation(dependency dependency.DependencyInterface, dependerLayer string, dependentLayer string) *SkippedViolation {
	return &SkippedViolation{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (v *SkippedViolation) GetDependency() dependency.DependencyInterface {
	return v.Dependency
}
func (v *SkippedViolation) GetDependerLayer() string {
	return v.DependerLayer
}
func (v *SkippedViolation) GetDependentLayer() string {
	return v.DependentLayer
}
