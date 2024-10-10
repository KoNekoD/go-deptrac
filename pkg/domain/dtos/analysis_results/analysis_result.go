package analysis_results

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/issues"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

// AnalysisResult - Describes the result_contract of a source code analysis.
type AnalysisResult struct {
	rules map[enums.RuleTypeEnum]map[string]violations_rules.RuleInterface

	warnings []*issues.Warning

	errors []*issues.Error
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		rules:    make(map[enums.RuleTypeEnum]map[string]violations_rules.RuleInterface),
		warnings: make([]*issues.Warning, 0),
		errors:   make([]*issues.Error, 0),
	}
}

func (r *AnalysisResult) ruleTypeByRule(rule violations_rules.RuleInterface) enums.RuleTypeEnum {
	switch rule.(type) {
	case *violations_rules.Violation:
		return enums.TypeViolation
	case *violations_rules.SkippedViolation:
		return enums.TypeSkippedViolation
	case *violations_rules.Uncovered:
		return enums.TypeUncovered
	case *violations_rules.Allowed:
		return enums.TypeAllowed
	default:
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}
}

func (r *AnalysisResult) AddRule(rule violations_rules.RuleInterface) {
	ruleType := r.ruleTypeByRule(rule)
	id := utils.SplObjectID(rule)

	if _, ok := r.rules[ruleType]; !ok {
		r.rules[ruleType] = make(map[string]violations_rules.RuleInterface)
	}

	r.rules[ruleType][id] = rule
}

func (r *AnalysisResult) RemoveRule(rule violations_rules.RuleInterface) {
	ruleType := r.ruleTypeByRule(rule)
	id := utils.SplObjectID(rule)

	delete(r.rules[ruleType], id)
}

func (r *AnalysisResult) Rules() map[enums.RuleTypeEnum]map[string]violations_rules.RuleInterface {
	return r.rules
}

func (r *AnalysisResult) AddWarning(warning *issues.Warning) {
	r.warnings = append(r.warnings, warning)
}

func (r *AnalysisResult) Warnings() []*issues.Warning {
	return r.warnings
}

func (r *AnalysisResult) AddError(error *issues.Error) {
	r.errors = append(r.errors, error)
}

func (r *AnalysisResult) Errors() []*issues.Error {
	return r.errors
}
