package enums

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

type RuleTypeEnum string

const (
	TypeViolation        RuleTypeEnum = "violation"
	TypeSkippedViolation RuleTypeEnum = "skipped_violation"
	TypeUncovered        RuleTypeEnum = "uncovered"
	TypeAllowed          RuleTypeEnum = "allowed"
)

func NewRuleTypeEnumByRule(rule rules.RuleInterface) RuleTypeEnum {
	switch rule.(type) {
	case *rules.Violation:
		return TypeViolation
	case *rules.SkippedViolation:
		return TypeSkippedViolation
	case *rules.Uncovered:
		return TypeUncovered
	case *rules.Allowed:
		return TypeAllowed
	default:
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}
}
