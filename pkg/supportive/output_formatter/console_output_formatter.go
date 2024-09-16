package output_formatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"strings"
)

type ConsoleOutputFormatter struct{}

func NewConsoleOutputFormatter() *ConsoleOutputFormatter {
	return &ConsoleOutputFormatter{}
}

func (f *ConsoleOutputFormatter) GetName() string {
	return "console"
}

func (f *ConsoleOutputFormatter) Finish(outputResult output_result.OutputResult, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) {
	for _, rule := range outputResult.AllOf(result.TypeViolation) {
		f.printViolation(rule.(*result.Violation), output)
	}

	if input.ReportSkipped {
		for _, rule := range outputResult.AllOf(result.TypeSkippedViolation) {
			f.printViolation(rule.(*result.SkippedViolation), output)
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

func (f *ConsoleOutputFormatter) printViolation(rule result.RuleInterface, output output_formatter.OutputInterface) {
	dep := rule.GetDependency()
	skippedText := ""

	dependerLayer := ""
	dependentLayer := ""

	if ruleAsserted, ok := rule.(*result.SkippedViolation); ok {
		skippedText = "[SKIPPED] "
		dependerLayer = ruleAsserted.GetDependerLayer()
		dependentLayer = ruleAsserted.GetDependentLayer()
	} else if ruleAsserted, ok := rule.(*result.Violation); ok {
		dependerLayer = ruleAsserted.GetDependerLayer()
		dependentLayer = ruleAsserted.GetDependentLayer()
	} else {
		panic(fmt.Errorf("unknown rule type: %T", rule))
	}

	output.WriteLineFormatted(
		output_formatter.StringOrArrayOfStrings{
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

func (f *ConsoleOutputFormatter) printMultilinePath(output output_formatter.OutputInterface, dep dependency.DependencyInterface) {
	var buffer strings.Builder
	for _, depSerialized := range dep.Serialize() {
		buffer.WriteString(fmt.Sprintf("\t%s:%d -> \n", depSerialized["name"], depSerialized["line"]))
	}
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: buffer.String()})
}

func (f *ConsoleOutputFormatter) printSummary(result output_result.OutputResult, output output_formatter.OutputInterface) {
	violationCount := len(result.Violations())
	skippedViolationCount := len(result.SkippedViolations())
	uncoveredCount := len(result.Uncovered())
	allowedCount := len(result.Allowed())
	warningsCount := len(result.Warnings)
	errorsCount := len(result.Errors)

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: "Report:"})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Violations: %d</>", f.getColor(violationCount > 0, "red", "default"), violationCount)})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Skipped violations: %d</>", f.getColor(skippedViolationCount > 0, "yellow", "default"), skippedViolationCount)})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Uncovered: %d</>", f.getColor(uncoveredCount > 0, "yellow", "default"), uncoveredCount)})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Allowed: %d</>", allowedCount)})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Warnings: %d</>", f.getColor(warningsCount > 0, "yellow", "default"), warningsCount)})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=%s>Errors: %d</>", f.getColor(errorsCount > 0, "red", "default"), errorsCount)})
}

func (f *ConsoleOutputFormatter) printUncovered(result output_result.OutputResult, output output_formatter.OutputInterface) {
	uncovered := result.Uncovered()
	if len(uncovered) == 0 {
		return
	}

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: "<comment>Uncovered dependencies:</>"})
	for _, u := range uncovered {
		dep := u.GetDependency()
		output.WriteLineFormatted(
			output_formatter.StringOrArrayOfStrings{
				String: fmt.Sprintf("<info>%s</> has uncovered dependency on <info>%s</> (%s)",
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

func (f *ConsoleOutputFormatter) printFileOccurrence(output output_formatter.OutputInterface, fileOccurrence *ast.FileOccurrence) {
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)})
}

func (f *ConsoleOutputFormatter) printErrors(result output_result.OutputResult, output output_formatter.OutputInterface) {
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
	for _, err := range result.Errors {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=red>[ERROR]</> %s", err)})
	}
}

func (f *ConsoleOutputFormatter) printWarnings(result output_result.OutputResult, output output_formatter.OutputInterface) {
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
	for _, warning := range result.Warnings {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<fg=yellow>[WARNING]</> %s", warning)})
	}
}

func (f *ConsoleOutputFormatter) getColor(condition bool, trueColor, falseColor string) string {
	if condition {
		return trueColor
	}
	return falseColor
}
