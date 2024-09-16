package result_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// RuleInterface - Represents a dependency_contract
type RuleInterface interface {
	GetDependency() dependency_contract.DependencyInterface
}
