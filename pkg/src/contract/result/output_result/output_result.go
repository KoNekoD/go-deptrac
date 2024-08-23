package output_result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/result"
)

// OutputResult - Represents a result ready for output formatting
type OutputResult struct {
	rules    map[result.RuleTypeEnum]map[string]result.RuleInterface
	Errors   []*result.Error
	Warnings []*result.Warning
}

func newOutputResult(rules map[result.RuleTypeEnum]map[string]result.RuleInterface, errors []*result.Error, warnings []*result.Warning) *OutputResult {
	return &OutputResult{rules: rules, Errors: errors, Warnings: warnings}
}

func NewOutputResultFromAnalysisResult(analysisResult *analysis_result.AnalysisResult) *OutputResult {
	return newOutputResult(
		analysisResult.Rules(),
		analysisResult.Errors(),
		analysisResult.Warnings(),
	)
}

func (r *OutputResult) AllOf(ruleType result.RuleTypeEnum) []result.RuleInterface {
	rulesByIds, ok := r.rules[ruleType]
	if ok {
		rules := make([]result.RuleInterface, 0, len(rulesByIds))

		for _, rule := range rulesByIds {
			rules = append(rules, rule)
		}

		return rules
	}
	return nil
}
func (r *OutputResult) AllRules() []result.RuleInterface {
	var rules []result.RuleInterface
	for _, ruleArray := range r.rules {
		for _, ruleItem := range ruleArray {
			rules = append(rules, ruleItem)
		}
	}
	return rules
}
func (r *OutputResult) Violations() []*result.Violation {
	untyped := r.AllOf(result.TypeViolation)

	items := make([]*result.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*result.SkippedViolation {
	untyped := r.AllOf(result.TypeSkippedViolation)

	items := make([]*result.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*result.Uncovered {
	untyped := r.AllOf(result.TypeUncovered)

	items := make([]*result.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result.Uncovered))
	}

	return items
}

func (r *OutputResult) HasUncovered() bool {
	return len(r.Uncovered()) > 0
}

func (r *OutputResult) HasAllowed() bool {
	return len(r.Allowed()) > 0
}

func (r *OutputResult) Allowed() []*result.Allowed {
	untyped := r.AllOf(result.TypeAllowed)

	items := make([]*result.Allowed, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result.Allowed))
	}

	return items
}

func (r *OutputResult) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *OutputResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}
