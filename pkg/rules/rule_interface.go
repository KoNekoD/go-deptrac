package rules

import "github.com/KoNekoD/go-deptrac/pkg/dependencies"

// RuleInterface - Represents a dependency_contract
type RuleInterface interface {
	GetDependency() dependencies.DependencyInterface
}
