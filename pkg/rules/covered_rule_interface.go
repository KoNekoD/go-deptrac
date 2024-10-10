package rules

// CoveredRuleInterface - Represents a dependency_contract that is covered by the defined rules. This does not mean that it is allowed to exist, just that it is covered. In that sense it exists as a complement to `Uncovered` struct
type CoveredRuleInterface interface {
	GetDependerLayer() string
	GetDependentLayer() string
}
