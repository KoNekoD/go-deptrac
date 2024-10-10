package formatters

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"strings"
)

type ConsoleOutputFormatter struct{}

func NewConsoleOutputFormatter() *ConsoleOutputFormatter {
	return &ConsoleOutputFormatter{}
}

func (f *ConsoleOutputFormatter) GetName() string {
	return "console_supportive"
}

func (f *ConsoleOutputFormatter) Finish(outputResult results.OutputResult, output results.OutputInterface, input OutputFormatterInput) {
	for _, rule := range outputResult.AllOf(enums.TypeViolation) {
		f.printViolation(rule.(*violations_rules.Violation), output)
	}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(enums.TypeSkippedViolation) {
			f.printViolation(rule.(*violations_rules.SkippedViolation), output)
		}
	}

	if input.ReportUncovered {
		f.printUncovered(outputResult, output)
	}

	if outputResult.HasErrors() {
		f.printErrors(outputResult, output)
	}

	if outputResult.HasWarnings() {
		f.printWarnings(outputResult, output)
	}

	f.printSummary(outputResult, output)
}

func (f *ConsoleOutputFormatter) printViolation(rule violations_rules.RuleInterface, output results.OutputInterface) {
	dep := rule.GetDependency()
	skippedText := ""

	dependerLayer := ""
	dependentLayer := ""

	if ruleAsserted, ok := rule.(*violations_rules.SkippedViolation); ok {
		skippedText = "[SKIPPED] "
		dependerLayer = ruleAsserted.GetDependerLayer()
		dependentLayer = ruleAsserted.GetDependentLayer()
	} else if ruleAsserted, ok := rule.(*violations_rules.Violation); ok {
		dependerLayer = ruleAsserted.GetDependerLayer()
		dependentLayer = ruleAsserted.GetDependentLayer()
	} else {
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}

	output.WriteLineFormatted(
		results.StringOrArrayOfStrings{
			String: fmt.Sprintf("%s<info>%s</> must not depend on <info>%s</> (%s on %s)",
				skippedText,
				dep.GetDepender().ToString(),
				dep.GetDependent().ToString(),
				dependerLayer,
				dependentLayer,
			),
		},
	)
	f.printFileOccurrence(output, dep.GetContext().FileOccurrence)

	if len(dep.Serialize()) > 1 {
		f.printMultilinePath(output, dep)
	}
}

func (f *ConsoleOutputFormatter) printMultilinePath(output results.OutputInterface, dep dependencies.DependencyInterface) {
	var buffer strings.Builder
	for _, depSerialized := range dep.Serialize() {
		buffer.WriteString(fmt.Sprintf("\t%s:%d -> \n", depSerialized["name"], depSerialized["line"]))
	}
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: buffer.String()})
}

func (f *ConsoleOutputFormatter) printSummary(result results.OutputResult, output results.OutputInterface) {
	violationCount := len(result.Violations())
	skippedViolationCount := len(result.SkippedViolations())
	uncoveredCount := len(result.Uncovered())
	allowedCount := len(result.Allowed())
	warningsCount := len(result.Warnings)
	errorsCount := len(result.Errors)

	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: "Report:"})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Violations: %d</>", f.getColor(violationCount > 0, "red", "default"), violationCount)})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Skipped violations: %d</>", f.getColor(skippedViolationCount > 0, "yellow", "default"), skippedViolationCount)})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Uncovered: %d</>", f.getColor(uncoveredCount > 0, "yellow", "default"), uncoveredCount)})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Allowed: %d</>", allowedCount)})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Warnings: %d</>", f.getColor(warningsCount > 0, "yellow", "default"), warningsCount)})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Errors: %d</>", f.getColor(errorsCount > 0, "red", "default"), errorsCount)})
}

func (f *ConsoleOutputFormatter) printUncovered(result results.OutputResult, output results.OutputInterface) {
	uncovered := result.Uncovered()
	if len(uncovered) == 0 {
		return
	}

	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: "<comment>Uncovered dependencies:</>"})
	for _, u := range uncovered {
		dep := u.GetDependency()
		output.WriteLineFormatted(
			results.StringOrArrayOfStrings{
				String: fmt.Sprintf("<info>%s</> has uncovered dependency_contract on <info>%s</> (%s)",
					dep.GetDepender().ToString(),
					dep.GetDependent().ToString(),
					u.Layer,
				),
			},
		)
		f.printFileOccurrence(output, dep.GetContext().FileOccurrence)

		if len(dep.Serialize()) > 1 {
			f.printMultilinePath(output, dep)
		}
	}
}

func (f *ConsoleOutputFormatter) printFileOccurrence(output results.OutputInterface, fileOccurrence *dtos.FileOccurrence) {
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)})
}

func (f *ConsoleOutputFormatter) printErrors(result results.OutputResult, output results.OutputInterface) {
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
	for _, err := range result.Errors {
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=red>[ERROR]</> %s", err)})
	}
}

func (f *ConsoleOutputFormatter) printWarnings(result results.OutputResult, output results.OutputInterface) {
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=yellow>[WARNING]</> %s", warning)})
	}
}

func (f *ConsoleOutputFormatter) getColor(condition bool, trueColor, falseColor string) string {
	if condition {
		return trueColor
	}
	return falseColor
}
