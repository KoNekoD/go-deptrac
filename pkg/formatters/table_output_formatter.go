package formatters

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	results2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/gookit/color"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

type TableOutputFormatter struct{}

func NewTableOutputFormatter() *TableOutputFormatter {
	return &TableOutputFormatter{}
}

func (t *TableOutputFormatter) GetName() enums.OutputFormatterType {
	return enums.Table
}

func (t *TableOutputFormatter) Finish(outputResult *results2.OutputResult, output pkg.OutputInterface, outputFormatterInput *OutputFormatterInput) error {
	groupedRules := make(map[string][]violations_rules.RuleInterface)

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
		rulesList := groupedRules[layer]

		slices.SortFunc(rulesList, func(a, b violations_rules.RuleInterface) int {
			if a.GetDependency().GetDepender().ToString() < b.GetDependency().GetDepender().ToString() {
				return -1
			}
			return 1
		})

		rows := make([][]string, 0)
		for _, ruleItem := range rulesList {
			switch item := ruleItem.(type) {
			case *violations_rules.Uncovered:
				rows = append(rows, t.uncoveredRow(item, outputFormatterInput.FailOnUncovered))
			case *violations_rules.Violation:
				rows = append(rows, t.violationRow(item))
			case *violations_rules.SkippedViolation:
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

func (t *TableOutputFormatter) skippedViolationRow(rule *violations_rules.SkippedViolation) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> must not depend on <info>%s</> (%s)", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString(), rule.GetDependentLayer())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)
	return []string{color.Sprint("<fg=yellow>Skipped</>"), message}

}

func (t *TableOutputFormatter) violationRow(rule *violations_rules.Violation) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> must not depend on <info>%s</>", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString())
	message += fmt.Sprintf("\n%s (To fix %s(You need to add to the array by this key) -> %s(That value needs to be added to that array))", rule.RuleDescription(), rule.GetDependerLayer(), rule.GetDependentLayer())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)
	return []string{color.Sprint("<fg=red>Violation</>"), message}
}
func (t *TableOutputFormatter) formatMultilinePath(dep dependencies.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " -> \n")
}

func (t *TableOutputFormatter) printSummary(result *results2.OutputResult, output pkg.OutputInterface, reportUncoveredAsError bool) {
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
		[]pkg.StringOrArrayOfStringsOrTableSeparator{
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

func (t *TableOutputFormatter) uncoveredRow(rule *violations_rules.Uncovered, reportAsError bool) []string {
	gotDependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> has uncovered dependency_contract on <info>%s</>", gotDependency.GetDepender().ToString(), gotDependency.GetDependent().ToString())
	if len(gotDependency.Serialize()) > 1 {
		message += "\n" + t.formatMultilinePath(gotDependency)
	}
	fileOccurrence := rule.GetDependency().GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)
	uncoveredFg := "yellow"
	if reportAsError {
		uncoveredFg = "red"
	}
	return []string{color.Sprintf("<fg=%s>Uncovered</>", uncoveredFg), message}
}

func (t *TableOutputFormatter) printErrors(result *results2.OutputResult, output pkg.OutputInterface) {
	errors := make([]string, 0)

	for _, e := range result.Errors {
		errors = append(errors, e.ToString())
	}

	output.GetStyle().Table([]string{color.Sprint("<fg=red>Errors</>")}, [][]string{errors})
}

func (t *TableOutputFormatter) printWarnings(result *results2.OutputResult, output pkg.OutputInterface) {
	warnings := make([]string, 0)

	for _, w := range result.Warnings {
		warnings = append(warnings, w.ToString())
	}

	output.GetStyle().Table([]string{color.Sprint("<fg=yellow>Warnings</>")}, [][]string{warnings})
}
