package result_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// SkippedViolation - Represents a Violation that is being skipped by the baseline file_supportive
type SkippedViolation struct {
	Dependency     dependency_contract.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewSkippedViolation(dependency dependency_contract.DependencyInterface, dependerLayer string, dependentLayer string) *SkippedViolation {
	return &SkippedViolation{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (v *SkippedViolation) GetDependency() dependency_contract.DependencyInterface {
	return v.Dependency
}
func (v *SkippedViolation) GetDependerLayer() string {
	return v.DependerLayer
}
func (v *SkippedViolation) GetDependentLayer() string {
	return v.DependentLayer
}
