package output_formatter_supportive

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract/output_result"
	"strings"
)

type GithubActionsOutputFormatter struct{}

func NewGithubActionsOutputFormatter() *GithubActionsOutputFormatter {
	return &GithubActionsOutputFormatter{}
}

func (g *GithubActionsOutputFormatter) GetName() output_formatter_contract2.OutputFormatterType {
	return output_formatter_contract2.GithubActions
}

func (g *GithubActionsOutputFormatter) Finish(outputResult *output_result.OutputResult, output output_formatter_contract2.OutputInterface, outputFormatterInput *output_formatter_contract2.OutputFormatterInput) error {
	for _, rule := range outputResult.AllOf(result_contract2.TypeViolation) {
		g.printViolation(rule, output)
	}
	if outputFormatterInput.ReportSkipped {
		for _, rule := range outputResult.AllOf(result_contract2.TypeSkippedViolation) {
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

func (g *GithubActionsOutputFormatter) determineLogLevel(rule result_contract2.RuleInterface) string {
	switch rule.(type) {
	case *result_contract2.Violation:
		return "error"
	case *result_contract2.SkippedViolation:
		return "warning"
	default:
		return "debug"
	}
}

func (g *GithubActionsOutputFormatter) printUncovered(result *output_result.OutputResult, output output_formatter_contract2.OutputInterface, reportAsError bool) {
	for _, u := range result.Uncovered() {
		dependency := u.GetDependency()

		reportAs := "warning"
		if reportAsError {
			reportAs = "error"
		}

		output.WriteLineFormatted(
			output_formatter_contract2.StringOrArrayOfStrings{
				String: fmt.Sprintf(
					"::%s file_supportive=%s,line=%d::%s has uncovered dependency_contract on %s (%s)",
					reportAs,
					dependency.GetContext().FileOccurrence.FilePath,
					dependency.GetContext().FileOccurrence.Line,
					dependency.GetDepender().ToString(),
					dependency.GetDependent().ToString(),
					u.Layer,
				),
			},
		)
	}
}

func (g *GithubActionsOutputFormatter) multilinePathMessage(dep dependency_contract.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " ->%0A")
}

func (g *GithubActionsOutputFormatter) printErrors(result *output_result.OutputResult, output output_formatter_contract2.OutputInterface) {
	for _, e := range result.Errors {
		output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("::error ::%s", e.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printWarnings(result *output_result.OutputResult, output output_formatter_contract2.OutputInterface) {
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("::warning ::%s", warning.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printViolation(rule result_contract2.RuleInterface, output output_formatter_contract2.OutputInterface) {
	dependency := rule.GetDependency()
	prefix := ""
	dependerLayer := ""
	dependentLayer := ""
	switch v := rule.(type) {
	case *result_contract2.SkippedViolation:
		prefix = "[SKIPPED] "
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	case *result_contract2.Violation:
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	}
	message := fmt.Sprintf("%s%s must not depend on %s (%s on %s)", prefix, dependency.GetDepender().ToString(), dependency.GetDependent().ToString(), dependerLayer, dependentLayer)
	if len(dependency.Serialize()) > 1 {
		message += "%0A" + g.multilinePathMessage(dependency)

	}
	output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("::%s file_supportive=%s,line=%d::%s", g.determineLogLevel(rule), dependency.GetContext().FileOccurrence.FilePath, dependency.GetContext().FileOccurrence.Line, message)})
}
