package output_formatter_supportive

import (
	"encoding/json"
	"fmt"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract/output_result"
	"os"
	"path/filepath"
)

type JsonOutputFormatter struct{}

func NewJsonOutputFormatter() *JsonOutputFormatter {
	return &JsonOutputFormatter{}
}

func (f *JsonOutputFormatter) GetName() string {
	return "json"
}

func (f *JsonOutputFormatter) Finish(outputResult output_result.OutputResult, output output_formatter_contract2.OutputInterface, input output_formatter_contract2.OutputFormatterInput) error {
	jsonArray := make(map[string]interface{})
	violations := make(map[string]FileViolations)

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(result_contract2.TypeSkippedViolation) {
			f.addSkipped(violations, rule.(*result_contract2.SkippedViolation))
		}
	}
	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(result_contract2.TypeUncovered) {
			f.addUncovered(violations, rule.(*result_contract2.Uncovered))
		}
	}
	for _, rule := range outputResult.AllOf(result_contract2.TypeViolation) {
		f.addFailure(violations, rule.(*result_contract2.Violation))
	}

	// Add report summary to jsonArray
	jsonArray["Report"] = map[string]interface{}{
		"Violations":         len(outputResult.Violations()),
		"Skipped violations": len(outputResult.SkippedViolations()),
		"Uncovered":          len(outputResult.Uncovered()),
		"Allowed":            len(outputResult.Allowed()),
		"Warnings":           len(outputResult.Warnings),
		"Errors":             len(outputResult.Errors),
	}

	// Add violation count to each file_supportive
	for fileName, fileViolation := range violations {
		fileViolation.Violations = len(fileViolation.Messages)
		violations[fileName] = fileViolation
	}

	jsonArray["files"] = violations

	jsonData, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to render JSON output: %v", err)
	}

	if input.OutputPath != nil && *input.OutputPath != "" {
		if err := os.WriteFile(*input.OutputPath, jsonData, 0644); err != nil {
			return err
		}
		output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("<info>JSON Report dumped to %s</>", filepath.Clean(*input.OutputPath))})
		return nil
	}

	output.WriteRaw(string(jsonData))
	return nil
}

func (f *JsonOutputFormatter) addFailure(violations map[string]FileViolations, violation *result_contract2.Violation) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getFailureMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "error",
	})
}

func (f *JsonOutputFormatter) getFailureMessage(violation *result_contract2.Violation) string {
	dependency := violation.GetDependency()
	return fmt.Sprintf("%s must not depend on %s (%s on %s)",
		dependency.GetDepender().ToString(),
		dependency.GetDependent().ToString(),
		violation.GetDependerLayer(),
		violation.GetDependentLayer(),
	)
}

func (f *JsonOutputFormatter) addSkipped(violations map[string]FileViolations, violation *result_contract2.SkippedViolation) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getWarningMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "warning",
	})
}

func (f *JsonOutputFormatter) getWarningMessage(violation *result_contract2.SkippedViolation) string {
	dependency := violation.GetDependency()
	return fmt.Sprintf("%s should not depend on %s (%s on %s)",
		dependency.GetDepender().ToString(),
		dependency.GetDependent().ToString(),
		violation.GetDependerLayer(),
		violation.GetDependentLayer(),
	)
}

func (f *JsonOutputFormatter) addUncovered(violations map[string]FileViolations, violation *result_contract2.Uncovered) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getUncoveredMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "warning",
	})
}

func (f *JsonOutputFormatter) getUncoveredMessage(violation *result_contract2.Uncovered) string {
	dependency := violation.GetDependency()
	return fmt.Sprintf("%s has uncovered dependency_contract on %s (%s)",
		dependency.GetDepender().ToString(),
		dependency.GetDependent().ToString(),
		violation.Layer,
	)
}

// Helper functions to manage violations
func appendViolation(violation FileViolations, message Message) FileViolations {
	violation.Messages = append(violation.Messages, message)
	return violation
}

type FileViolations struct {
	Messages   []Message `json:"messages"`
	Violations int       `json:"violations"`
}

type Message struct {
	Message string `json:"message"`
	Line    int    `json:"line"`
	Type    string `json:"type"`
}
