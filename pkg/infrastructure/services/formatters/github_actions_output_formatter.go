package formatters

import (
	"fmt"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"strings"
)

type GithubActionsOutputFormatter struct{}

func NewGithubActionsOutputFormatter() *GithubActionsOutputFormatter {
	return &GithubActionsOutputFormatter{}
}

func (g *GithubActionsOutputFormatter) GetName() enums.OutputFormatterType {
	return enums.GithubActions
}

func (g *GithubActionsOutputFormatter) Finish(outputResult *results.OutputResult, output services2.OutputInterface, outputFormatterInput *OutputFormatterInput) error {
	for _, rule := range outputResult.AllOf(enums.TypeViolation) {
		g.printViolation(rule, output)
	}
	if outputFormatterInput.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums.TypeSkippedViolation) {
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

func (g *GithubActionsOutputFormatter) determineLogLevel(rule violations_rules.RuleInterface) string {
	switch rule.(type) {
	case *violations_rules.Violation:
		return "error"
	case *violations_rules.SkippedViolation:
		return "warning"
	default:
		return "debug"
	}
}

func (g *GithubActionsOutputFormatter) printUncovered(result *results.OutputResult, output services2.OutputInterface, reportAsError bool) {
	for _, u := range result.Uncovered() {
		dependency := u.GetDependency()

		reportAs := "warning"
		if reportAsError {
			reportAs = "error"
		}

		output.WriteLineFormatted(
			services2.StringOrArrayOfStrings{
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

func (g *GithubActionsOutputFormatter) multilinePathMessage(dep dependencies.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " ->%0A")
}

func (g *GithubActionsOutputFormatter) printErrors(result *results.OutputResult, output services2.OutputInterface) {
	for _, e := range result.Errors {
		output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: fmt.Sprintf("::error ::%s", e.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printWarnings(result *results.OutputResult, output services2.OutputInterface) {
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: fmt.Sprintf("::warning ::%s", warning.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printViolation(rule violations_rules.RuleInterface, output services2.OutputInterface) {
	dependency := rule.GetDependency()
	prefix := ""
	dependerLayer := ""
	dependentLayer := ""
	switch v := rule.(type) {
	case *violations_rules.SkippedViolation:
		prefix = "[SKIPPED] "
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	case *violations_rules.Violation:
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	}
	message := fmt.Sprintf("%s%s must not depend on %s (%s on %s)", prefix, dependency.GetDepender().ToString(), dependency.GetDependent().ToString(), dependerLayer, dependentLayer)
	if len(dependency.Serialize()) > 1 {
		message += "%0A" + g.multilinePathMessage(dependency)

	}
	output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: fmt.Sprintf("::%s file_supportive=%s,line=%d::%s", g.determineLogLevel(rule), dependency.GetContext().FileOccurrence.FilePath, dependency.GetContext().FileOccurrence.Line, message)})
}
