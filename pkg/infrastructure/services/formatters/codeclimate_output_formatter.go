package formatters

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"os"
)

type CodeclimateOutputFormatter struct {
	config map[enums.SeverityType]interface{}
}

func NewCodeclimateOutputFormatter(config FormatterConfiguration) *CodeclimateOutputFormatter {
	extractedConfig := config.GetConfigFor("codeclimate").(interface{}).(map[enums.SeverityType]interface{})
	return &CodeclimateOutputFormatter{config: extractedConfig}
}

func (f *CodeclimateOutputFormatter) GetName() string {
	return "codeclimate"
}

func (f *CodeclimateOutputFormatter) Finish(outputResult results.OutputResult, output services.OutputInterface, input OutputFormatterInput) error {
	formatterConfig := enums.NewConfigurationCodeclimateFromArray(f.config)
	var violationsList []map[string]interface{}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums.TypeSkippedViolation) {
			f.addSkipped(&violationsList, rule.(*violations_rules.SkippedViolation), formatterConfig)
		}
	}

	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(enums.TypeUncovered) {
			f.addUncovered(&violationsList, rule.(*violations_rules.Uncovered), formatterConfig)
		}
	}

	for _, rule := range outputResult.AllOf(enums.TypeViolation) {
		f.addFailure(&violationsList, rule.(*violations_rules.Violation), formatterConfig)
	}

	jsonData, err := json.MarshalIndent(violationsList, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to render codeclimate output: %v", err)
	}

	if input.OutputPath != nil && *input.OutputPath != "" {
		err := os.WriteFile(*input.OutputPath, jsonData, 0644)
		if err != nil {
			return err
		}
		output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Codeclimate Report dumped to %s</>", *input.OutputPath)})
		return nil
	}

	output.WriteRaw(string(jsonData))
	return nil
}

func (f *CodeclimateOutputFormatter) addFailure(violations *[]map[string]interface{}, violation *violations_rules.Violation, config *enums.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getFailureMessage(violation), config.GetSeverity("failure")))
}

func (f *CodeclimateOutputFormatter) getFailureMessage(violation *violations_rules.Violation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s must not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addSkipped(violations *[]map[string]interface{}, violation *violations_rules.SkippedViolation, config *enums.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getWarningMessage(violation), config.GetSeverity("skipped")))
}

func (f *CodeclimateOutputFormatter) getWarningMessage(violation *violations_rules.SkippedViolation) *string {
	dependency := violation.GetDependency()
	return utils.AsPtr(fmt.Sprintf("%s should not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addUncovered(violations *[]map[string]interface{}, violation *violations_rules.Uncovered, config *enums.ConfigurationCodeclimate) {
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
