package analysis_result

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Error"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleTypeEnum"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Warning"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// AnalysisResult - Describes the result of a source code analysis.
type AnalysisResult struct {
	rules map[RuleTypeEnum.RuleTypeEnum]map[string]RuleInterface.RuleInterface

	warnings []*Warning.Warning

	errors []*Error.Error
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		rules:    make(map[RuleTypeEnum.RuleTypeEnum]map[string]RuleInterface.RuleInterface),
		warnings: make([]*Warning.Warning, 0),
		errors:   make([]*Error.Error, 0),
	}
}

func (r *AnalysisResult) AddRule(rule RuleInterface.RuleInterface) {
	ruleType := RuleTypeEnum.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	if _, ok := r.rules[ruleType]; !ok {
		r.rules[ruleType] = make(map[string]RuleInterface.RuleInterface)
	}

	r.rules[ruleType][id] = rule
}

func (r *AnalysisResult) RemoveRule(rule RuleInterface.RuleInterface) {
	ruleType := RuleTypeEnum.NewRuleTypeEnumByRule(rule)
	id := util.SplObjectID(rule)

	delete(r.rules[ruleType], id)
}

func (r *AnalysisResult) Rules() map[RuleTypeEnum.RuleTypeEnum]map[string]RuleInterface.RuleInterface {
	return r.rules
}

func (r *AnalysisResult) AddWarning(warning *Warning.Warning) {
	r.warnings = append(r.warnings, warning)
}

func (r *AnalysisResult) Warnings() []*Warning.Warning {
	return r.warnings
}

func (r *AnalysisResult) AddError(error *Error.Error) {
	r.errors = append(r.errors, error)
}

func (r *AnalysisResult) Errors() []*Error.Error {
	return r.errors
}
