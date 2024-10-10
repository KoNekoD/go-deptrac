package results

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

// OutputResult - Represents a result_contract ready for output formatting
type OutputResult struct {
	rules    map[enums.RuleTypeEnum]map[string]rules.RuleInterface
	Errors   []*apperrors.Error
	Warnings []*tokens.Warning
}

func newOutputResult(rules map[enums.RuleTypeEnum]map[string]rules.RuleInterface, errors []*apperrors.Error, warnings []*tokens.Warning) *OutputResult {
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
func (r *OutputResult) Violations() []*rules.Violation {
	untyped := r.AllOf(enums.TypeViolation)

	items := make([]*rules.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*rules.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*rules.SkippedViolation {
	untyped := r.AllOf(enums.TypeSkippedViolation)

	items := make([]*rules.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*rules.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*rules.Uncovered {
	untyped := r.AllOf(enums.TypeUncovered)

	items := make([]*rules.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*rules.Uncovered))
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
