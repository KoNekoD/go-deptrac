package RuleTypeEnum

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Allowed"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/SkippedViolation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Violation"
)

type RuleTypeEnum string

const (
	TypeViolation        RuleTypeEnum = "violation"
	TypeSkippedViolation RuleTypeEnum = "skipped_violation"
	TypeUncovered        RuleTypeEnum = "uncovered"
	TypeAllowed          RuleTypeEnum = "allowed"
)

func NewRuleTypeEnumByRule(rule RuleInterface.RuleInterface) RuleTypeEnum {
	switch rule.(type) {
	case *Violation.Violation:
		return TypeViolation
	case *SkippedViolation.SkippedViolation:
		return TypeSkippedViolation
	case *Uncovered.Uncovered:
		return TypeUncovered
	case *Allowed.Allowed:
		return TypeAllowed
	default:
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}
}
