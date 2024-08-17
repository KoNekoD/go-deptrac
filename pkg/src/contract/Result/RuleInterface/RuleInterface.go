package RuleInterface

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"

// RuleInterface - Represents a dependency
type RuleInterface interface {
	GetDependency() DependencyInterface.DependencyInterface
}
