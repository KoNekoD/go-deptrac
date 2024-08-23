package AnalyseRunner

import (
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInput"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/OutputResult"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseOptions"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/CommandRunException"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/FormatterProvider"
	"github.com/hashicorp/go-multierror"
	"strings"
)

// AnalyseRunner - Should only be used by AnalyseCommand
type AnalyseRunner struct {
	analyzer          *analyser.DependencyLayersAnalyser
	formatterProvider *FormatterProvider.FormatterProvider
}

func NewAnalyseRunner(analyzer *analyser.DependencyLayersAnalyser, formatterProvider *FormatterProvider.FormatterProvider) *AnalyseRunner {
	return &AnalyseRunner{
		analyzer:          analyzer,
		formatterProvider: formatterProvider,
	}
}

func (r *AnalyseRunner) Run(options *AnalyseOptions.AnalyseOptions, output OutputInterface.OutputInterface) error {
	outputFormatterType, err := OutputFormatterType.NewOutputFormatterTypeFromString(options.Formatter)
	if err != nil {
		return err
	}
	formatter, err := r.formatterProvider.Get(outputFormatterType)
	if err != nil {
		r.printFormatterNotFoundException(output, options.Formatter)
		return CommandRunException.NewCommandRunExceptionInvalidFormatter()
	}
	formatterInput := OutputFormatterInput.NewOutputFormatterInput(*options.Output, options.ReportSkipped, options.ReportUncovered, options.FailOnUncovered)
	r.printCollectViolations(output)

	analysisResult, errAnalyse := r.analyzer.Analyse()
	if errAnalyse != nil {
		r.printAnalysisException(output, multierror.Append(errAnalyse))
		return CommandRunException.NewCommandRunExceptionAnalyserException(errAnalyse)
	}
	result := OutputResult.NewOutputResultFromAnalysisResult(analysisResult)
	r.printFormattingStart(output)
	errFinish := formatter.Finish(result, output, formatterInput)
	if errFinish != nil {
		r.printFormatterError(output, string(formatter.GetName()), errFinish)
	}
	if options.FailOnUncovered && result.HasUncovered() {
		return CommandRunException.NewCommandRunExceptionFinishedWithUncovered()
	}
	if result.HasViolations() {
		return CommandRunException.NewCommandRunExceptionFinishedWithViolations()
	}
	if result.HasErrors() {
		return CommandRunException.NewCommandRunExceptionFailedWithErrors()
	}

	return nil
}

func (r *AnalyseRunner) printCollectViolations(output OutputInterface.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: "<info>collecting violations.</>"})
	}
}

func (r *AnalyseRunner) printFormattingStart(output OutputInterface.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: "<info>formatting dependencies.</>"})
	}
}

func (r *AnalyseRunner) printFormatterError(output OutputInterface.OutputInterface, formatterName string, error error) {
	output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: ""})
	output.GetStyle().Error(OutputStyleInterface.StringOrArrayOfStrings{Strings: []string{"", fmt.Sprintf("OutputInterface formatter %s threw an Exception:", formatterName), fmt.Sprintf("Message: %s", error.Error()), ""}})
	output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: ""})
}

var JsonMultiErrFormatFunc = func(es []error) string {
	errorsStrings := make([]string, len(es))
	for i, err := range es {
		errorsStrings[i] = err.Error()
	}

	marshalled, err := json.Marshal(errorsStrings)

	if err != nil {
		return "(marshall json err) " + err.Error()
	}

	return string(marshalled)
}

func (r *AnalyseRunner) printAnalysisException(output OutputInterface.OutputInterface, exception *multierror.Error) {
	message := []string{"Analysis finished with an Exception.", JsonMultiErrFormatFunc(exception.Errors), ""}
	output.GetStyle().Error(OutputStyleInterface.StringOrArrayOfStrings{Strings: message})
}

func (r *AnalyseRunner) printFormatterNotFoundException(output OutputInterface.OutputInterface, formatterName string) {
	output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: ""})

	knownFormatters := make([]string, 0)
	for _, formatterType := range r.formatterProvider.GetKnownFormatters() {
		knownFormatters = append(knownFormatters, string(formatterType))
	}

	output.GetStyle().Error(OutputStyleInterface.StringOrArrayOfStrings{Strings: []string{fmt.Sprintf("Output formatter %s not found.", formatterName), "Available formatters:", strings.Join(knownFormatters, ", "), ""}})
	output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: ""})
}
