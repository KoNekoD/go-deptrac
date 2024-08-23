package result

import (
	"fmt"
)

type RuleTypeEnum string

const (
	TypeViolation        RuleTypeEnum = "violation"
	TypeSkippedViolation RuleTypeEnum = "skipped_violation"
	TypeUncovered        RuleTypeEnum = "uncovered"
	TypeAllowed          RuleTypeEnum = "allowed"
)

func NewRuleTypeEnumByRule(rule RuleInterface) RuleTypeEnum {
	switch rule.(type) {
	case *Violation:
		return TypeViolation
	case *SkippedViolation:
		return TypeSkippedViolation
	case *Uncovered:
		return TypeUncovered
	case *Allowed:
		return TypeAllowed
	default:
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}
}
