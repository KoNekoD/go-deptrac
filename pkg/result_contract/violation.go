package result_contract

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/violation_creating_interface"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// Violation - Represents a dependency_contract that is NOT allowed to exist given the defined rules
type Violation struct {
	Dependency            dependency_contract.DependencyInterface
	DependerLayer         string
	DependentLayer        string
	ViolationCreatingRule violation_creating_interface.ViolationCreatingInterface
}

func NewViolation(dependency dependency_contract.DependencyInterface, dependerLayer string, dependentLayer string, violationCreatingRule violation_creating_interface.ViolationCreatingInterface) *Violation {

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

func (v *Violation) GetDependency() dependency_contract.DependencyInterface {
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
