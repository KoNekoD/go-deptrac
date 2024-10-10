package rules

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/violations"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

// AnalysisResult - Describes the result_contract of a source code analysis.
type AnalysisResult struct {
	rules map[enums.RuleTypeEnum]map[string]RuleInterface

	warnings []*violations.Warning

	errors []*violations.Error
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		rules:    make(map[enums.RuleTypeEnum]map[string]RuleInterface),
		warnings: make([]*violations.Warning, 0),
		errors:   make([]*violations.Error, 0),
	}
}

func (r *AnalysisResult) ruleTypeByRule(rule RuleInterface) enums.RuleTypeEnum {
	switch rule.(type) {
	case *violations.Violation:
		return enums.TypeViolation
	case *violations.SkippedViolation:
		return enums.TypeSkippedViolation
	case *violations.Uncovered:
		return enums.TypeUncovered
	case *Allowed:
		return enums.TypeAllowed
	default:
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}
}

func (r *AnalysisResult) AddRule(rule RuleInterface) {
	ruleType := r.ruleTypeByRule(rule)
	id := utils.SplObjectID(rule)

	if _, ok := r.rules[ruleType]; !ok {
		r.rules[ruleType] = make(map[string]RuleInterface)
	}

	r.rules[ruleType][id] = rule
}

func (r *AnalysisResult) RemoveRule(rule RuleInterface) {
	ruleType := r.ruleTypeByRule(rule)
	id := utils.SplObjectID(rule)

	delete(r.rules[ruleType], id)
}

func (r *AnalysisResult) Rules() map[enums.RuleTypeEnum]map[string]RuleInterface {
	return r.rules
}

func (r *AnalysisResult) AddWarning(warning *violations.Warning) {
	r.warnings = append(r.warnings, warning)
}

func (r *AnalysisResult) Warnings() []*violations.Warning {
	return r.warnings
}

func (r *AnalysisResult) AddError(error *violations.Error) {
	r.errors = append(r.errors, error)
}

func (r *AnalysisResult) Errors() []*violations.Error {
	return r.errors
}
