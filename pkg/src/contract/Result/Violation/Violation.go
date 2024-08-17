package Violation

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ViolationCreatingInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
)

// Violation - Represents a dependency that is NOT allowed to exist given the defined rules
type Violation struct {
	Dependency            DependencyInterface.DependencyInterface
	DependerLayer         string
	DependentLayer        string
	ViolationCreatingRule ViolationCreatingInterface.ViolationCreatingInterface
}

func NewViolation(dependency DependencyInterface.DependencyInterface, dependerLayer string, dependentLayer string, violationCreatingRule ViolationCreatingInterface.ViolationCreatingInterface) *Violation {

	if dependentLayer == dependerLayer {
		panic("1")
	}

	if dependerLayer == "File" && dependentLayer == "Ast" {
		fmt.Println()
	}

	return &Violation{
		Dependency:            dependency,
		DependerLayer:         dependerLayer,
		DependentLayer:        dependentLayer,
		ViolationCreatingRule: violationCreatingRule,
	}
}

func (v *Violation) GetDependency() DependencyInterface.DependencyInterface {
	return v.Dependency
}
func (v *Violation) GetDependerLayer() string {
	return v.DependerLayer
}
func (v *Violation) GetDependentLayer() string {
	return v.DependentLayer
}
func (v *Violation) RuleName() string {
	return v.ViolationCreatingRule.RuleName()
}

func (v *Violation) RuleDescription() string {
	return v.ViolationCreatingRule.RuleDescription()
}
