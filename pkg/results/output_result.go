package results

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results"
	analysis_results_failure2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/issues"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

// OutputResult - Represents a result_contract ready for output formatting
type OutputResult struct {
	rules    map[enums.RuleTypeEnum]map[string]violations_rules.RuleInterface
	Errors   []*analysis_results_failure2.Error
	Warnings []*analysis_results_failure2.Warning
}

func newOutputResult(rules map[enums.RuleTypeEnum]map[string]violations_rules.RuleInterface, errors []*analysis_results_failure2.Error, warnings []*analysis_results_failure2.Warning) *OutputResult {
	return &OutputResult{rules: rules, Errors: errors, Warnings: warnings}
}

func NewOutputResultFromAnalysisResult(analysisResult *analysis_results.AnalysisResult) *OutputResult {
	return newOutputResult(
		analysisResult.Rules(),
		analysisResult.Errors(),
		analysisResult.Warnings(),
	)
}

func (r *OutputResult) AllOf(ruleType enums.RuleTypeEnum) []violations_rules.RuleInterface {
	rulesByIds, ok := r.rules[ruleType]
	if ok {
		rules := make([]violations_rules.RuleInterface, 0, len(rulesByIds))

		for _, rule := range rulesByIds {
			rules = append(rules, rule)
		}

		return rules
	}
	return nil
}
func (r *OutputResult) AllRules() []violations_rules.RuleInterface {
	var rules []violations_rules.RuleInterface
	for _, ruleArray := range r.rules {
		for _, ruleItem := range ruleArray {
			rules = append(rules, ruleItem)
		}
	}
	return rules
}
func (r *OutputResult) Violations() []*violations_rules.Violation {
	untyped := r.AllOf(enums.TypeViolation)

	items := make([]*violations_rules.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations_rules.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*violations_rules.SkippedViolation {
	untyped := r.AllOf(enums.TypeSkippedViolation)

	items := make([]*violations_rules.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations_rules.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*violations_rules.Uncovered {
	untyped := r.AllOf(enums.TypeUncovered)

	items := make([]*violations_rules.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations_rules.Uncovered))
	}

	return items
}

func (r *OutputResult) HasUncovered() bool {
	return len(r.Uncovered()) > 0
}

func (r *OutputResult) HasAllowed() bool {
	return len(r.Allowed()) > 0
}

func (r *OutputResult) Allowed() []*violations_rules.Allowed {
	untyped := r.AllOf(enums.TypeAllowed)

	items := make([]*violations_rules.Allowed, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*violations_rules.Allowed))
	}

	return items
}

func (r *OutputResult) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *OutputResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}
