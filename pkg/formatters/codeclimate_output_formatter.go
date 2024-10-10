package formatters

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	results2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"os"
)

type CodeclimateOutputFormatter struct {
	config map[enums2.SeverityType]interface{}
}

func NewCodeclimateOutputFormatter(config FormatterConfiguration) *CodeclimateOutputFormatter {
	extractedConfig := config.GetConfigFor("codeclimate").(interface{}).(map[enums2.SeverityType]interface{})
	return &CodeclimateOutputFormatter{config: extractedConfig}
}

func (f *CodeclimateOutputFormatter) GetName() string {
	return "codeclimate"
}

func (f *CodeclimateOutputFormatter) Finish(outputResult results2.OutputResult, output results.OutputInterface, input OutputFormatterInput) error {
	formatterConfig := enums2.NewConfigurationCodeclimateFromArray(f.config)
	var violations []map[string]interface{}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums2.TypeSkippedViolation) {
			f.addSkipped(&violations, rule.(*violations.SkippedViolation), formatterConfig)
		}
	}

	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(enums2.TypeUncovered) {
			f.addUncovered(&violations, rule.(*violations.Uncovered), formatterConfig)
		}
	}

	for _, rule := range outputResult.AllOf(enums2.TypeViolation) {
		f.addFailure(&violations, rule.(*violations.Violation), formatterConfig)
	}

	jsonData, err := json.MarshalIndent(violations, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to render codeclimate output: %v", err)
	}

	if input.OutputPath != nil && *input.OutputPath != "" {
		err := os.WriteFile(*input.OutputPath, jsonData, 0644)
		if err != nil {
			return err
		}
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Codeclimate Report dumped to %s</>", *input.OutputPath)})
		return nil
	}

	output.WriteRaw(string(jsonData))
	return nil
}

func (f *CodeclimateOutputFormatter) addFailure(violations *[]map[string]interface{}, violation *violations_rules.Violation, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getFailureMessage(violation), config.GetSeverity("failure")))
}

func (f *CodeclimateOutputFormatter) getFailureMessage(violation *violations_rules.Violation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s must not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addSkipped(violations *[]map[string]interface{}, violation *violations_rules.SkippedViolation, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getWarningMessage(violation), config.GetSeverity("skipped")))
}

func (f *CodeclimateOutputFormatter) getWarningMessage(violation *violations_rules.SkippedViolation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s should not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addUncovered(violations *[]map[string]interface{}, violation *violations_rules.Uncovered, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getUncoveredMessage(violation), config.GetSeverity("uncovered")))
}

func (f *CodeclimateOutputFormatter) getUncoveredMessage(violation *violations_rules.Uncovered) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s has uncovered dependency_contract on %s (%s)", dependency.GetDepender(), dependency.GetDependent(), violation.Layer))
}

func (f *CodeclimateOutputFormatter) buildRuleArray(rule violations_rules.RuleInterface, message, severity *string) map[string]interface{} {
	return map[string]interface{}{
		"type":        "issue",
		"check_name":  "Dependency violation",
		"fingerprint": f.buildFingerprint(rule),
		"description": message,
		"categories":  []string{"Style", "Complexity"},
		"severity":    severity,
		"location": map[string]interface{}{
			"path": rule.GetDependency().GetContext().FileOccurrence.FilePath,
			"lines": map[string]interface{}{
				"begin": rule.GetDependency().GetContext().FileOccurrence.Line,
			},
		},
	}
}

func (f *CodeclimateOutputFormatter) buildFingerprint(rule violations_rules.RuleInterface) string {
	data := fmt.Sprintf("%s,%s,%s,%s,%d",
		rule,
		rule.GetDependency().GetDepender(),
		rule.GetDependency().GetDependent(),
		rule.GetDependency().GetContext().FileOccurrence.FilePath,
		rule.GetDependency().GetContext().FileOccurrence.Line)

	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
