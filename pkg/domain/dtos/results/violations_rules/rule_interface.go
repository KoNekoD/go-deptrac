package violations_rules

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
)

// RuleInterface - Represents a dependency_contract
type RuleInterface interface {
	GetDependency() dependencies.DependencyInterface
}
