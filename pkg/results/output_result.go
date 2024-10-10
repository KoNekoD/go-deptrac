package results

import (
	violations2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/violations"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

// OutputResult - Represents a result_contract ready for output formatting
type OutputResult struct {
	rules    map[enums.RuleTypeEnum]map[string]rules.RuleInterface
	Errors   []*violations2.Error
	Warnings []*violations2.Warning
}

func newOutputResult(rules map[enums.RuleTypeEnum]map[string]rules.RuleInterface, errors []*violations2.Error, warnings []*violations2.Warning) *OutputResult {
	return &OutputResult{rules: rules, Errors: errors, Warnings: warnings}
}

func NewOutputResultFromAnalysisResult(analysisResult *rules.AnalysisResult) *OutputResult {
	return newOutputResult(
		analysisResult.Rules(),
		analysisResult.Errors(),
		analysisResult.Warnings(),
	)
}

func (r *OutputResult) AllOf(ruleType enums.RuleTypeEnum) []rules.RuleInterface {
	rulesByIds, ok := r.rules[ruleType]
	if ok {
		rules := make([]rules.RuleInterface, 0, len(rulesByIds))

		for _, rule := range rulesByIds {
			rules = append(rules, rule)
		}

		return rules
	}
	return nil
}
func (r *OutputResult) AllRules() []rules.RuleInterface {
	var rules []rules.RuleInterface
	for _, ruleArray := range r.rules {
		for _, ruleItem := range ruleArray {
			rules = append(rules, ruleItem)
		}
	}
	return rules
}
func (r *OutputResult) Violations() []*violations2.Violation {
	untyped := r.AllOf(enums.TypeViolation)

	items := make([]*violations2.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations2.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*violations2.SkippedViolation {
	untyped := r.AllOf(enums.TypeSkippedViolation)

	items := make([]*violations2.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations2.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*violations2.Uncovered {
	untyped := r.AllOf(enums.TypeUncovered)

	items := make([]*violations2.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations2.Uncovered))
	}

	return items
}

func (r *OutputResult) HasUncovered() bool {
	return len(r.Uncovered()) > 0
}

func (r *OutputResult) HasAllowed() bool {
	return len(r.Allowed()) > 0
}

func (r *OutputResult) Allowed() []*rules.Allowed {
	untyped := r.AllOf(enums.TypeAllowed)

	items := make([]*rules.Allowed, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*rules.Allowed))
	}

	return items
}

func (r *OutputResult) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *OutputResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}
