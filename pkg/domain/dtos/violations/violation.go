package violations

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

// Violation - Represents a dependency_contract that is NOT allowed to exist given the defined rules
type Violation struct {
	Dependency            dependencies.DependencyInterface
	DependerLayer         string
	DependentLayer        string
	ViolationCreatingRule ViolationCreatingInterface
}

func NewViolation(dependency dependencies.DependencyInterface, dependerLayer string, dependentLayer string, violationCreatingRule ViolationCreatingInterface) *Violation {

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

func (v *Violation) GetDependency() dependencies.DependencyInterface {
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
