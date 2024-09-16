package analysis_result

import (
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// AnalysisResult - Describes the result_contract of a source code analysis.
type AnalysisResult struct {
	rules map[result_contract2.RuleTypeEnum]map[string]result_contract2.RuleInterface

	warnings []*result_contract2.Warning

	errors []*result_contract2.Error
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		rules:    make(map[result_contract2.RuleTypeEnum]map[string]result_contract2.RuleInterface),
		warnings: make([]*result_contract2.Warning, 0),
		errors:   make([]*result_contract2.Error, 0),
	}
}

func (r *AnalysisResult) AddRule(rule result_contract2.RuleInterface) {
	ruleType := result_contract2.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	if _, ok := r.rules[ruleType]; !ok {
		r.rules[ruleType] = make(map[string]result_contract2.RuleInterface)
	}

	r.rules[ruleType][id] = rule
}

func (r *AnalysisResult) RemoveRule(rule result_contract2.RuleInterface) {
	ruleType := result_contract2.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	delete(r.rules[ruleType], id)
}

func (r *AnalysisResult) Rules() map[result_contract2.RuleTypeEnum]map[string]result_contract2.RuleInterface {
	return r.rules
}

func (r *AnalysisResult) AddWarning(warning *result_contract2.Warning) {
	r.warnings = append(r.warnings, warning)
}

func (r *AnalysisResult) Warnings() []*result_contract2.Warning {
	return r.warnings
}

func (r *AnalysisResult) AddError(error *result_contract2.Error) {
	r.errors = append(r.errors, error)
}

func (r *AnalysisResult) Errors() []*result_contract2.Error {
	return r.errors
}
