package result

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/dependency"
)

// RuleInterface - Represents a dependency
type RuleInterface interface {
	GetDependency() dependency.DependencyInterface
}
