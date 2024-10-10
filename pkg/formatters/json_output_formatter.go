package formatters

import (
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg"
	results2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
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

func (f *JsonOutputFormatter) Finish(outputResult results2.OutputResult, output pkg.OutputInterface, input OutputFormatterInput) error {
	jsonArray := make(map[string]interface{})
	violations := make(map[string]FileViolations)

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums.TypeSkippedViolation) {
			f.addSkipped(violations, rule.(*violations.SkippedViolation))
		}
	}
	if input.ReportUncovered {
		for _, rule := range outputResult.AllOf(enums.TypeUncovered) {
			f.addUncovered(violations, rule.(*violations.Uncovered))
		}
	}
	for _, rule := range outputResult.AllOf(enums.TypeViolation) {
		f.addFailure(violations, rule.(*violations.Violation))
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
		output.WriteLineFormatted(pkg.StringOrArrayOfStrings{String: fmt.Sprintf("<info>JSON Report dumped to %s</>", filepath.Clean(*input.OutputPath))})
		return nil
	}

	output.WriteRaw(string(jsonData))
	return nil
}

func (f *JsonOutputFormatter) addFailure(violations map[string]FileViolations, violation *violations_rules.Violation) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getFailureMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "error",
	})
}

func (f *JsonOutputFormatter) getFailureMessage(violation *violations_rules.Violation) string {
	dependency := violation.GetDependency()
	return fmt.Sprintf("%s must not depend on %s (%s on %s)",
		dependency.GetDepender().ToString(),
		dependency.GetDependent().ToString(),
		violation.GetDependerLayer(),
		violation.GetDependentLayer(),
	)
}

func (f *JsonOutputFormatter) addSkipped(violations map[string]FileViolations, violation *violations_rules.SkippedViolation) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getWarningMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "warning",
	})
}

func (f *JsonOutputFormatter) getWarningMessage(violation *violations_rules.SkippedViolation) string {
	dependency := violation.GetDependency()
	return fmt.Sprintf("%s should not depend on %s (%s on %s)",
		dependency.GetDepender().ToString(),
		dependency.GetDependent().ToString(),
		violation.GetDependerLayer(),
		violation.GetDependentLayer(),
	)
}

func (f *JsonOutputFormatter) addUncovered(violations map[string]FileViolations, violation *violations_rules.Uncovered) {
	className := violation.GetDependency().GetContext().FileOccurrence.FilePath
	violations[className] = appendViolation(violations[className], Message{
		Message: f.getUncoveredMessage(violation),
		Line:    violation.GetDependency().GetContext().FileOccurrence.Line,
		Type:    "warning",
	})
}

func (f *JsonOutputFormatter) getUncoveredMessage(violation *violations_rules.Uncovered) string {
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
