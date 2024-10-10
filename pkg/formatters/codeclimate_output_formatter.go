package formatters

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
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

func (f *CodeclimateOutputFormatter) Finish(outputResult results.OutputResult, output results.OutputInterface, input OutputFormatterInput) error {
	formatterConfig := enums2.NewConfigurationCodeclimateFromArray(f.config)
	var violations []map[string]interface{}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums2.TypeSkippedViolation) {
			f.addSkipped(&violations, rule.(*rules.SkippedViolation), formatterConfig)
		}
	}

	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(enums2.TypeUncovered) {
			f.addUncovered(&violations, rule.(*rules.Uncovered), formatterConfig)
		}
	}

	for _, rule := range outputResult.AllOf(enums2.TypeViolation) {
		f.addFailure(&violations, rule.(*rules.Violation), formatterConfig)
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

func (f *CodeclimateOutputFormatter) addFailure(violations *[]map[string]interface{}, violation *rules.Violation, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getFailureMessage(violation), config.GetSeverity("failure")))
}

func (f *CodeclimateOutputFormatter) getFailureMessage(violation *rules.Violation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s must not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addSkipped(violations *[]map[string]interface{}, violation *rules.SkippedViolation, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getWarningMessage(violation), config.GetSeverity("skipped")))
}

func (f *CodeclimateOutputFormatter) getWarningMessage(violation *rules.SkippedViolation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s should not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addUncovered(violations *[]map[string]interface{}, violation *rules.Uncovered, config *enums2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getUncoveredMessage(violation), config.GetSeverity("uncovered")))
}

func (f *CodeclimateOutputFormatter) getUncoveredMessage(violation *rules.Uncovered) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s has uncovered dependency_contract on %s (%s)", dependency.GetDepender(), dependency.GetDependent(), violation.Layer))
}

func (f *CodeclimateOutputFormatter) buildRuleArray(rule rules.RuleInterface, message, severity *string) map[string]interface{} {
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

func (f *CodeclimateOutputFormatter) buildFingerprint(rule rules.RuleInterface) string {
	data := fmt.Sprintf("%s,%s,%s,%s,%d",
		rule,
		rule.GetDependency().GetDepender(),
		rule.GetDependency().GetDependent(),
		rule.GetDependency().GetContext().FileOccurrence.FilePath,
		rule.GetDependency().GetContext().FileOccurrence.Line)

	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
