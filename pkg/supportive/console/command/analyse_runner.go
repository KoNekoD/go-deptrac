package command

import (
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"github.com/KoNekoD/go-deptrac/pkg/core/analyser"
	output_formatter2 "github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter"
	"github.com/hashicorp/go-multierror"
	"strings"
)

// AnalyseRunner - Should only be used by AnalyseCommand
type AnalyseRunner struct {
	analyzer          *analyser.DependencyLayersAnalyser
	formatterProvider *output_formatter2.FormatterProvider
}

func NewAnalyseRunner(analyzer *analyser.DependencyLayersAnalyser, formatterProvider *output_formatter2.FormatterProvider) *AnalyseRunner {
	return &AnalyseRunner{
		analyzer:          analyzer,
		formatterProvider: formatterProvider,
	}
}

func (r *AnalyseRunner) Run(options *AnalyseOptions, output output_formatter.OutputInterface) error {
	outputFormatterType, err := output_formatter.NewOutputFormatterTypeFromString(options.Formatter)
	if err != nil {
		return err
	}
	formatter, err := r.formatterProvider.Get(outputFormatterType)
	if err != nil {
		r.printFormatterNotFoundException(output, options.Formatter)
		return NewCommandRunExceptionInvalidFormatter()
	}
	formatterInput := output_formatter.NewOutputFormatterInput(*options.Output, options.ReportSkipped, options.ReportUncovered, options.FailOnUncovered)
	r.printCollectViolations(output)

	analysisResult, errAnalyse := r.analyzer.Analyse()
	if errAnalyse != nil {
		r.printAnalysisException(output, multierror.Append(errAnalyse))
		return NewCommandRunExceptionAnalyserException(errAnalyse)
	}
	result := output_result.NewOutputResultFromAnalysisResult(analysisResult)
	r.printFormattingStart(output)
	errFinish := formatter.Finish(result, output, formatterInput)
	if errFinish != nil {
		r.printFormatterError(output, string(formatter.GetName()), errFinish)
	}
	if options.FailOnUncovered && result.HasUncovered() {
		return NewCommandRunExceptionFinishedWithUncovered()
	}
	if result.HasViolations() {
		return NewCommandRunExceptionFinishedWithViolations()
	}
	if result.HasErrors() {
		return NewCommandRunExceptionFailedWithErrors()
	}

	return nil
}

func (r *AnalyseRunner) printCollectViolations(output output_formatter.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: "<info>collecting violations.</>"})
	}
}

func (r *AnalyseRunner) printFormattingStart(output output_formatter.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: "<info>formatting dependencies.</>"})
	}
}

func (r *AnalyseRunner) printFormatterError(output output_formatter.OutputInterface, formatterName string, error error) {
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
	output.GetStyle().Error(output_formatter.StringOrArrayOfStrings{Strings: []string{"", fmt.Sprintf("OutputInterface formatter %s threw an Exception:", formatterName), fmt.Sprintf("Message: %s", error.Error()), ""}})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
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

func (r *AnalyseRunner) printAnalysisException(output output_formatter.OutputInterface, exception *multierror.Error) {
	message := []string{"Analysis finished with an Exception.", JsonMultiErrFormatFunc(exception.Errors), ""}
	output.GetStyle().Error(output_formatter.StringOrArrayOfStrings{Strings: message})
}

func (r *AnalyseRunner) printFormatterNotFoundException(output output_formatter.OutputInterface, formatterName string) {
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})

	knownFormatters := make([]string, 0)
	for _, formatterType := range r.formatterProvider.GetKnownFormatters() {
		knownFormatters = append(knownFormatters, string(formatterType))
	}

	output.GetStyle().Error(output_formatter.StringOrArrayOfStrings{Strings: []string{fmt.Sprintf("Output formatter %s not found.", formatterName), "Available formatters:", strings.Join(knownFormatters, ", "), ""}})
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: ""})
}
