package OutputResult

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Allowed"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Error"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleTypeEnum"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/SkippedViolation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Violation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Warning"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/analysis_result"
)

// OutputResult - Represents a result ready for output formatting
type OutputResult struct {
	rules    map[RuleTypeEnum.RuleTypeEnum]map[string]RuleInterface.RuleInterface
	Errors   []*Error.Error
	Warnings []*Warning.Warning
}

func newOutputResult(rules map[RuleTypeEnum.RuleTypeEnum]map[string]RuleInterface.RuleInterface, errors []*Error.Error, warnings []*Warning.Warning) *OutputResult {
	return &OutputResult{rules: rules, Errors: errors, Warnings: warnings}
}

func NewOutputResultFromAnalysisResult(analysisResult *analysis_result.AnalysisResult) *OutputResult {
	return newOutputResult(
		analysisResult.Rules(),
		analysisResult.Errors(),
		analysisResult.Warnings(),
	)
}

func (r *OutputResult) AllOf(ruleType RuleTypeEnum.RuleTypeEnum) []RuleInterface.RuleInterface {
	rulesByIds, ok := r.rules[ruleType]
	if ok {
		rules := make([]RuleInterface.RuleInterface, 0, len(rulesByIds))

		for _, rule := range rulesByIds {
			rules = append(rules, rule)
		}

		return rules
	}
	return nil
}
func (r *OutputResult) AllRules() []RuleInterface.RuleInterface {
	var rules []RuleInterface.RuleInterface
	for _, ruleArray := range r.rules {
		for _, ruleItem := range ruleArray {
			rules = append(rules, ruleItem)
		}
	}
	return rules
}
func (r *OutputResult) Violations() []*Violation.Violation {
	untyped := r.AllOf(RuleTypeEnum.TypeViolation)

	items := make([]*Violation.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*Violation.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*SkippedViolation.SkippedViolation {
	untyped := r.AllOf(RuleTypeEnum.TypeSkippedViolation)

	items := make([]*SkippedViolation.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*SkippedViolation.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*Uncovered.Uncovered {
	untyped := r.AllOf(RuleTypeEnum.TypeUncovered)

	items := make([]*Uncovered.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*Uncovered.Uncovered))
	}

	return items
}

func (r *OutputResult) HasUncovered() bool {
	return len(r.Uncovered()) > 0
}

func (r *OutputResult) HasAllowed() bool {
	return len(r.Allowed()) > 0
}

func (r *OutputResult) Allowed() []*Allowed.Allowed {
	untyped := r.AllOf(RuleTypeEnum.TypeAllowed)

	items := make([]*Allowed.Allowed, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*Allowed.Allowed))
	}

	return items
}

func (r *OutputResult) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *OutputResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}
