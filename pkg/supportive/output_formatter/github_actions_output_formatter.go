package output_formatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"strings"
)

type GithubActionsOutputFormatter struct{}

func NewGithubActionsOutputFormatter() *GithubActionsOutputFormatter {
	return &GithubActionsOutputFormatter{}
}

func (g *GithubActionsOutputFormatter) GetName() output_formatter.OutputFormatterType {
	return output_formatter.GithubActions
}

func (g *GithubActionsOutputFormatter) Finish(outputResult *output_result.OutputResult, output output_formatter.OutputInterface, outputFormatterInput *output_formatter.OutputFormatterInput) error {
	for _, rule := range outputResult.AllOf(result.TypeViolation) {
		g.printViolation(rule, output)
	}
	if outputFormatterInput.ReportSkipped {
		for _, rule := range outputResult.AllOf(result.TypeSkippedViolation) {
			g.printViolation(rule, output)
		}
	}
	if outputFormatterInput.ReportUncovered {
		g.printUncovered(outputResult, output, outputFormatterInput.FailOnUncovered)
	}
	if outputResult.HasErrors() {
		g.printErrors(outputResult, output)
	}
	if outputResult.HasWarnings() {
		g.printWarnings(outputResult, output)
	}

	return nil
}

func (g *GithubActionsOutputFormatter) determineLogLevel(rule result.RuleInterface) string {
	switch rule.(type) {
	case *result.Violation:
		return "error"
	case *result.SkippedViolation:
		return "warning"
	default:
		return "debug"
	}
}

func (g *GithubActionsOutputFormatter) printUncovered(result *output_result.OutputResult, output output_formatter.OutputInterface, reportAsError bool) {
	for _, u := range result.Uncovered() {
		dependency := u.GetDependency()

		reportAs := "warning"
		if reportAsError {
			reportAs = "error"
		}

		output.WriteLineFormatted(
			output_formatter.StringOrArrayOfStrings{
				String: fmt.Sprintf(
					"::%s file=%s,line=%d::%s has uncovered dependency on %s (%s)",
					reportAs,
					dependency.GetContext().FileOccurrence.Filepath,
					dependency.GetContext().FileOccurrence.Line,
					dependency.GetDepender().ToString(),
					dependency.GetDependent().ToString(),
					u.Layer,
				),
			},
		)
	}
}

func (g *GithubActionsOutputFormatter) multilinePathMessage(dep dependency.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " ->%0A")
}

func (g *GithubActionsOutputFormatter) printErrors(result *output_result.OutputResult, output output_formatter.OutputInterface) {
	for _, e := range result.Errors {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("::error ::%s", e.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printWarnings(result *output_result.OutputResult, output output_formatter.OutputInterface) {
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("::warning ::%s", warning.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printViolation(rule result.RuleInterface, output output_formatter.OutputInterface) {
	dependency := rule.GetDependency()
	prefix := ""
	dependerLayer := ""
	dependentLayer := ""
	switch v := rule.(type) {
	case *result.SkippedViolation:
		prefix = "[SKIPPED] "
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	case *result.Violation:
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	}
	message := fmt.Sprintf("%s%s must not depend on %s (%s on %s)", prefix, dependency.GetDepender().ToString(), dependency.GetDependent().ToString(), dependerLayer, dependentLayer)
	if len(dependency.Serialize()) > 1 {
		message += "%0A" + g.multilinePathMessage(dependency)

	}
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("::%s file=%s,line=%d::%s", g.determineLogLevel(rule), dependency.GetContext().FileOccurrence.Filepath, dependency.GetContext().FileOccurrence.Line, message)})
}
