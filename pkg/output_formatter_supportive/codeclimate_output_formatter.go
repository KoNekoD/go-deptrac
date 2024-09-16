package output_formatter_supportive

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	configuration2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive/configuration"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract/output_result"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
)

type CodeclimateOutputFormatter struct {
	config map[configuration2.SeverityType]interface{}
}

func NewCodeclimateOutputFormatter(config configuration2.FormatterConfiguration) *CodeclimateOutputFormatter {
	extractedConfig := config.GetConfigFor("codeclimate").(interface{}).(map[configuration2.SeverityType]interface{})
	return &CodeclimateOutputFormatter{config: extractedConfig}
}

func (f *CodeclimateOutputFormatter) GetName() string {
	return "codeclimate"
}

func (f *CodeclimateOutputFormatter) Finish(outputResult output_result.OutputResult, output output_formatter_contract2.OutputInterface, input output_formatter_contract2.OutputFormatterInput) error {
	formatterConfig := configuration2.NewConfigurationCodeclimateFromArray(f.config)
	var violations []map[string]interface{}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(result_contract2.TypeSkippedViolation) {
			f.addSkipped(&violations, rule.(*result_contract2.SkippedViolation), formatterConfig)
		}
	}

	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(result_contract2.TypeUncovered) {
			f.addUncovered(&violations, rule.(*result_contract2.Uncovered), formatterConfig)
		}
	}

	for _, rule := range outputResult.AllOf(result_contract2.TypeViolation) {
		f.addFailure(&violations, rule.(*result_contract2.Violation), formatterConfig)
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
		output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Codeclimate Report dumped to %s</>", *input.OutputPath)})
		return nil
	}

	output.WriteRaw(string(jsonData))
	return nil
}

func (f *CodeclimateOutputFormatter) addFailure(violations *[]map[string]interface{}, violation *result_contract2.Violation, config *configuration2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getFailureMessage(violation), config.GetSeverity("failure")))
}

func (f *CodeclimateOutputFormatter) getFailureMessage(violation *result_contract2.Violation) *string {
	dependency := violation.GetDependency()
	return util.AsPtr(fmt.Sprintf("%s must not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addSkipped(violations *[]map[string]interface{}, violation *result_contract2.SkippedViolation, config *configuration2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getWarningMessage(violation), config.GetSeverity("skipped")))
}

func (f *CodeclimateOutputFormatter) getWarningMessage(violation *result_contract2.SkippedViolation) *string {
	dependency := violation.GetDependency()
	return util.AsPtr(fmt.Sprintf("%s should not depend on %s (%s on %s)", dependency.GetDepender(), dependency.GetDependent(), violation.GetDependerLayer(), violation.GetDependentLayer()))
}

func (f *CodeclimateOutputFormatter) addUncovered(violations *[]map[string]interface{}, violation *result_contract2.Uncovered, config *configuration2.ConfigurationCodeclimate) {
	*violations = append(*violations, f.buildRuleArray(violation, f.getUncoveredMessage(violation), config.GetSeverity("uncovered")))
}

func (f *CodeclimateOutputFormatter) getUncoveredMessage(violation *result_contract2.Uncovered) *string {
	dependency := violation.GetDependency()
	return util.AsPtr(fmt.Sprintf("%s has uncovered dependency_contract on %s (%s)", dependency.GetDepender(), dependency.GetDependent(), violation.Layer))
}

func (f *CodeclimateOutputFormatter) buildRuleArray(rule result_contract2.RuleInterface, message, severity *string) map[string]interface{} {
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

func (f *CodeclimateOutputFormatter) buildFingerprint(rule result_contract2.RuleInterface) string {
	data := fmt.Sprintf("%s,%s,%s,%s,%d",
		rule,
		rule.GetDependency().GetDepender(),
		rule.GetDependency().GetDependent(),
		rule.GetDependency().GetContext().FileOccurrence.FilePath,
		rule.GetDependency().GetContext().FileOccurrence.Line)

	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
