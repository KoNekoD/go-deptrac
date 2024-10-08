package TableOutputFormatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInput"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/OutputResult"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/SkippedViolation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Violation"
	"github.com/gookit/color"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

type TableOutputFormatter struct{}

func NewTableOutputFormatter() *TableOutputFormatter {
	return &TableOutputFormatter{}
}

func (t *TableOutputFormatter) GetName() OutputFormatterType.OutputFormatterType {
	return OutputFormatterType.Table
}

func (t *TableOutputFormatter) Finish(outputResult *OutputResult.OutputResult, output OutputInterface.OutputInterface, outputFormatterInput *OutputFormatterInput.OutputFormatterInput) error {
	groupedRules := make(map[string][]RuleInterface.RuleInterface)

	for _, ruleItem := range outputResult.Violations() {
		groupedRules[ruleItem.GetDependerLayer()] = append(groupedRules[ruleItem.GetDependerLayer()], ruleItem)
	}

	if outputFormatterInput.ReportSkipped {
		for _, ruleItem := range outputResult.SkippedViolations() {
			groupedRules[ruleItem.GetDependerLayer()] = append(groupedRules[ruleItem.GetDependerLayer()], ruleItem)
		}
	}

	if outputFormatterInput.ReportUncovered {
		for _, ruleItem := range outputResult.Uncovered() {
			groupedRules[ruleItem.Layer] = append(groupedRules[ruleItem.Layer], ruleItem)
		}
	}

	groupedRulesLayers := maps.Keys(groupedRules)
	slices.Sort(groupedRulesLayers)

	style := output.GetStyle()
	for _, layer := range groupedRulesLayers {
		rules := groupedRules[layer]

		slices.SortFunc(rules, func(a, b RuleInterface.RuleInterface) int {
			if a.GetDependency().GetDepender().ToString() < b.GetDependency().GetDepender().ToString() {
				return -1
			}
			return 1
		})

		rows := make([][]string, 0)
		for _, ruleItem := range rules {
			switch item := ruleItem.(type) {
			case *Uncovered.Uncovered:
				rows = append(rows, t.uncoveredRow(item, outputFormatterInput.FailOnUncovered))
			case *Violation.Violation:
				rows = append(rows, t.violationRow(item))
			case *SkippedViolation.SkippedViolation:
				rows = append(rows, t.skippedViolationRow(item))
			}
		}

		style.Table([]string{"Reason", layer}, rows)
	}
	if outputResult.HasErrors() {
		t.printErrors(outputResult, output)
	}
	if outputResult.HasWarnings() {
		t.printWarnings(outputResult, output)
	}
	t.printSummary(outputResult, output, outputFormatterInput.FailOnUncovered)
	return nil
}

func (t *TableOutputFormatter) skippedViolationRow(rule *SkippedViolation.SkippedViolation) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> must not depend on <info>%s</> (%s)", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString(), rule.GetDependentLayer())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.Filepath, fileOccurrence.Line)
	return []string{color.Sprint("<fg=yellow>Skipped</>"), message}

}

func (t *TableOutputFormatter) violationRow(rule *Violation.Violation) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> must not depend on <info>%s</>", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString())
	message += fmt.Sprintf("\n%s (%s -> %s)", rule.RuleDescription(), rule.GetDependerLayer(), rule.GetDependentLayer())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.Filepath, fileOccurrence.Line)
	return []string{color.Sprint("<fg=red>Violation</>"), message}
}
func (t *TableOutputFormatter) formatMultilinePath(dep DependencyInterface.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " -> \n")
}

func (t *TableOutputFormatter) printSummary(result *OutputResult.OutputResult, output OutputInterface.OutputInterface, reportUncoveredAsError bool) {
	violationCount := len(result.Violations())
	skippedViolationCount := len(result.SkippedViolations())
	uncoveredCount := len(result.Uncovered())
	allowedCount := len(result.Allowed())
	warningsCount := len(result.Warnings)
	errorsCount := len(result.Errors)

	uncoveredFg := "red"
	if !reportUncoveredAsError {
		uncoveredFg = "yellow"
	}

	violationsColor := "default"
	if violationCount > 0 {
		violationsColor = "red"
	}

	skippedViolationsColor := "default"
	if skippedViolationCount > 0 {
		skippedViolationsColor = "yellow"
	}

	uncoveredColor := "default"
	if uncoveredCount > 0 {
		uncoveredColor = uncoveredFg
	}

	warningColor := "default"
	if warningsCount > 0 {
		warningColor = "yellow"
	}

	errorColor := "default"
	if errorsCount > 0 {
		errorColor = "red"
	}

	style := output.GetStyle()
	style.NewLine(1)
	style.DefinitionList(
		[]OutputStyleInterface.StringOrArrayOfStringsOrTableSeparator{
			{String: "Report"},
			{TableSeparator: true},
			{StringsMap: map[string]string{"Violations": color.Sprintf("<fg=%s>%d</>", violationsColor, violationCount)}},
			{StringsMap: map[string]string{"Skipped violations": color.Sprintf("<fg=%s>%d</>", skippedViolationsColor, skippedViolationCount)}},
			{StringsMap: map[string]string{"Uncovered": color.Sprintf("<fg=%s>%d</>", uncoveredColor, uncoveredCount)}},
			{StringsMap: map[string]string{"Allowed": color.Sprintf("%d", allowedCount)}},
			{StringsMap: map[string]string{"Warnings": color.Sprintf("<fg=%s>%d</>", warningColor, warningsCount)}},
			{StringsMap: map[string]string{"Errors": color.Sprintf("<fg=%s>%d</>", errorColor, errorsCount)}},
		},
	)
}

func (t *TableOutputFormatter) uncoveredRow(rule *Uncovered.Uncovered, reportAsError bool) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> has uncovered dependency on <info>%s</>", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.Filepath, fileOccurrence.Line)
	uncoveredFg := "yellow"
	if reportAsError {
		uncoveredFg = "red"
	}
	return []string{color.Sprintf("<fg=%s>Uncovered</>", uncoveredFg), message}
}

func (t *TableOutputFormatter) printErrors(result *OutputResult.OutputResult, output OutputInterface.OutputInterface) {
	errors := make([]string, 0)

	for _, e := range result.Errors {
		errors = append(errors, e.ToString())
	}

	output.GetStyle().Table([]string{color.Sprint("<fg=red>Errors</>")}, [][]string{errors})
}

func (t *TableOutputFormatter) printWarnings(result *OutputResult.OutputResult, output OutputInterface.OutputInterface) {
	warnings := make([]string, 0)

	for _, w := range result.Warnings {
		warnings = append(warnings, w.ToString())
	}

	output.GetStyle().Table([]string{color.Sprint("<fg=yellow>Warnings</>")}, [][]string{warnings})
}
