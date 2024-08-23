package analysis_result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// AnalysisResult - Describes the result of a source code analysis.
type AnalysisResult struct {
	rules map[result.RuleTypeEnum]map[string]result.RuleInterface

	warnings []*result.Warning

	errors []*result.Error
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		rules:    make(map[result.RuleTypeEnum]map[string]result.RuleInterface),
		warnings: make([]*result.Warning, 0),
		errors:   make([]*result.Error, 0),
	}
}

func (r *AnalysisResult) AddRule(rule result.RuleInterface) {
	ruleType := result.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	if _, ok := r.rules[ruleType]; !ok {
		r.rules[ruleType] = make(map[string]result.RuleInterface)
	}

	r.rules[ruleType][id] = rule
}

func (r *AnalysisResult) RemoveRule(rule result.RuleInterface) {
	ruleType := result.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	delete(r.rules[ruleType], id)
}

func (r *AnalysisResult) Rules() map[result.RuleTypeEnum]map[string]result.RuleInterface {
	return r.rules
}

func (r *AnalysisResult) AddWarning(warning *result.Warning) {
	r.warnings = append(r.warnings, warning)
}

func (r *AnalysisResult) Warnings() []*result.Warning {
	return r.warnings
}

func (r *AnalysisResult) AddError(error *result.Error) {
	r.errors = append(r.errors, error)
}

func (r *AnalysisResult) Errors() []*result.Error {
	return r.errors
}
