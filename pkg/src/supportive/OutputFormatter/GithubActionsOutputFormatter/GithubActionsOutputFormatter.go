package GithubActionsOutputFormatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInput"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/OutputResult"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/RuleTypeEnum"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/SkippedViolation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Violation"
	"strings"
)

type GithubActionsOutputFormatter struct{}

func NewGithubActionsOutputFormatter() *GithubActionsOutputFormatter {
	return &GithubActionsOutputFormatter{}
}

func (g *GithubActionsOutputFormatter) GetName() OutputFormatterType.OutputFormatterType {
	return OutputFormatterType.GithubActions
}

func (g *GithubActionsOutputFormatter) Finish(result *OutputResult.OutputResult, output OutputInterface.OutputInterface, outputFormatterInput *OutputFormatterInput.OutputFormatterInput) error {
	for _, rule := range result.AllOf(RuleTypeEnum.TypeViolation) {
		g.printViolation(rule, output)
	}
	if outputFormatterInput.ReportSkipped {
		for _, rule := range result.AllOf(RuleTypeEnum.TypeSkippedViolation) {
			g.printViolation(rule, output)
		}
	}
	if outputFormatterInput.ReportUncovered {
		g.printUncovered(result, output, outputFormatterInput.FailOnUncovered)
	}
	if result.HasErrors() {
		g.printErrors(result, output)
	}
	if result.HasWarnings() {
		g.printWarnings(result, output)
	}

	return nil
}

func (g *GithubActionsOutputFormatter) determineLogLevel(rule RuleInterface.RuleInterface) string {
	switch rule.(type) {
	case *Violation.Violation:
		return "error"
	case *SkippedViolation.SkippedViolation:
		return "warning"
	default:
		return "debug"
	}
}

func (g *GithubActionsOutputFormatter) printUncovered(result *OutputResult.OutputResult, output OutputInterface.OutputInterface, reportAsError bool) {
	for _, u := range result.Uncovered() {
		dependency := u.GetDependency()

		reportAs := "warning"
		if reportAsError {
			reportAs = "error"
		}

		output.WriteLineFormatted(
			OutputStyleInterface.StringOrArrayOfStrings{
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

func (g *GithubActionsOutputFormatter) multilinePathMessage(dep DependencyInterface.DependencyInterface) string {
	lines := make([]string, 0)
	for _, serializedDependency := range dep.Serialize() {
		lines = append(lines, fmt.Sprintf("%s::%d", serializedDependency["name"], serializedDependency["line"]))
	}
	return strings.Join(lines, " ->%0A")
}

func (g *GithubActionsOutputFormatter) printErrors(result *OutputResult.OutputResult, output OutputInterface.OutputInterface) {
	for _, e := range result.Errors {
		output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("::error ::%s", e.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printWarnings(result *OutputResult.OutputResult, output OutputInterface.OutputInterface) {
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("::warning ::%s", warning.ToString())})
	}
}

func (g *GithubActionsOutputFormatter) printViolation(rule RuleInterface.RuleInterface, output OutputInterface.OutputInterface) {
	dependency := rule.GetDependency()
	prefix := ""
	dependerLayer := ""
	dependentLayer := ""
	switch v := rule.(type) {
	case *SkippedViolation.SkippedViolation:
		prefix = "[SKIPPED] "
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	case *Violation.Violation:
		dependerLayer = v.GetDependerLayer()
		dependentLayer = v.GetDependentLayer()
	}
	message := fmt.Sprintf("%s%s must not depend on %s (%s on %s)", prefix, dependency.GetDepender().ToString(), dependency.GetDependent().ToString(), dependerLayer, dependentLayer)
	if len(dependency.Serialize()) > 1 {
		message += "%0A" + g.multilinePathMessage(dependency)

	}
	output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("::%s file=%s,line=%d::%s", g.determineLogLevel(rule), dependency.GetContext().FileOccurrence.Filepath, dependency.GetContext().FileOccurrence.Line, message)})
}
