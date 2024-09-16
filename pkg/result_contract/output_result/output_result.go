package output_result

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/analysis_result"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
)

// OutputResult - Represents a result_contract ready for output formatting
type OutputResult struct {
	rules    map[result_contract2.RuleTypeEnum]map[string]result_contract2.RuleInterface
	Errors   []*result_contract2.Error
	Warnings []*result_contract2.Warning
}

func newOutputResult(rules map[result_contract2.RuleTypeEnum]map[string]result_contract2.RuleInterface, errors []*result_contract2.Error, warnings []*result_contract2.Warning) *OutputResult {
	return &OutputResult{rules: rules, Errors: errors, Warnings: warnings}
}

func NewOutputResultFromAnalysisResult(analysisResult *analysis_result.AnalysisResult) *OutputResult {
	return newOutputResult(
		analysisResult.Rules(),
		analysisResult.Errors(),
		analysisResult.Warnings(),
	)
}

func (r *OutputResult) AllOf(ruleType result_contract2.RuleTypeEnum) []result_contract2.RuleInterface {
	rulesByIds, ok := r.rules[ruleType]
	if ok {
		rules := make([]result_contract2.RuleInterface, 0, len(rulesByIds))

		for _, rule := range rulesByIds {
			rules = append(rules, rule)
		}

		return rules
	}
	return nil
}
func (r *OutputResult) AllRules() []result_contract2.RuleInterface {
	var rules []result_contract2.RuleInterface
	for _, ruleArray := range r.rules {
		for _, ruleItem := range ruleArray {
			rules = append(rules, ruleItem)
		}
	}
	return rules
}
func (r *OutputResult) Violations() []*result_contract2.Violation {
	untyped := r.AllOf(result_contract2.TypeViolation)

	items := make([]*result_contract2.Violation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result_contract2.Violation))
	}

	return items
}

func (r *OutputResult) HasViolations() bool {
	return len(r.Violations()) > 0
}

func (r *OutputResult) SkippedViolations() []*result_contract2.SkippedViolation {
	untyped := r.AllOf(result_contract2.TypeSkippedViolation)

	items := make([]*result_contract2.SkippedViolation, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result_contract2.SkippedViolation))
	}

	return items
}

func (r *OutputResult) Uncovered() []*result_contract2.Uncovered {
	untyped := r.AllOf(result_contract2.TypeUncovered)

	items := make([]*result_contract2.Uncovered, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result_contract2.Uncovered))
	}

	return items
}

func (r *OutputResult) HasUncovered() bool {
	return len(r.Uncovered()) > 0
}

func (r *OutputResult) HasAllowed() bool {
	return len(r.Allowed()) > 0
}

func (r *OutputResult) Allowed() []*result_contract2.Allowed {
	untyped := r.AllOf(result_contract2.TypeAllowed)

	items := make([]*result_contract2.Allowed, 0, len(untyped))
	for _, item := range untyped {
		items = append(items, item.(*result_contract2.Allowed))
	}

	return items
}

func (r *OutputResult) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *OutputResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}
