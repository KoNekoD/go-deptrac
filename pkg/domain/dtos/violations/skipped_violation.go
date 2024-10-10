package violations

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

// SkippedViolation - Represents a Violation that is being skipped by the baseline file_supportive
type SkippedViolation struct {
	Dependency     dependencies.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewSkippedViolation(dependency dependencies.DependencyInterface, dependerLayer string, dependentLayer string) *SkippedViolation {
	return &SkippedViolation{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (v *SkippedViolation) GetDependency() dependencies.DependencyInterface {
	return v.Dependency
}
func (v *SkippedViolation) GetDependerLayer() string {
	return v.DependerLayer
}
func (v *SkippedViolation) GetDependentLayer() string {
	return v.DependentLayer
}
