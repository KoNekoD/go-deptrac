package SkippedViolation

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"

// SkippedViolation - Represents a Violation that is being skipped by the baseline file
type SkippedViolation struct {
	Dependency     DependencyInterface.DependencyInterface
	DependerLayer  string
	DependentLayer string
}

func NewSkippedViolation(dependency DependencyInterface.DependencyInterface, dependerLayer string, dependentLayer string) *SkippedViolation {
	return &SkippedViolation{
		Dependency:     dependency,
		DependerLayer:  dependerLayer,
		DependentLayer: dependentLayer,
	}
}

func (v *SkippedViolation) GetDependency() DependencyInterface.DependencyInterface {
	return v.Dependency
}
func (v *SkippedViolation) GetDependerLayer() string {
	return v.DependerLayer
}
func (v *SkippedViolation) GetDependentLayer() string {
	return v.DependentLayer
}
