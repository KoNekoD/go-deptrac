package app

import (
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
	"github.com/hashicorp/go-multierror"
	"strings"
)

// AnalyseRunner - Should only be used by AnalyseCommand
type AnalyseRunner struct {
	analyzer          *subscribers.DependencyLayersAnalyser
	formatterProvider *formatters.FormatterProvider
}

func NewAnalyseRunner(analyzer *subscribers.DependencyLayersAnalyser, formatterProvider *formatters.FormatterProvider) *AnalyseRunner {
	return &AnalyseRunner{
		analyzer:          analyzer,
		formatterProvider: formatterProvider,
	}
}

func (r *AnalyseRunner) Run(options *rules.AnalyseOptions, output results.OutputInterface) error {
	outputFormatterType, err := enums.NewOutputFormatterTypeFromString(options.Formatter)
	if err != nil {
		return err
	}
	formatter, err := r.formatterProvider.Get(outputFormatterType)
	if err != nil {
		r.printFormatterNotFoundException(output, options.Formatter)
		return apperrors.NewCommandRunExceptionInvalidFormatter()
	}
	formatterInput := formatters.NewOutputFormatterInput(*options.Output, options.ReportSkipped, options.ReportUncovered, options.FailOnUncovered)
	r.printCollectViolations(output)

	analysisResult, errAnalyse := r.analyzer.Analyse()
	if errAnalyse != nil {
		r.printAnalysisException(output, multierror.Append(errAnalyse))
		return apperrors.NewCommandRunExceptionAnalyserException(errAnalyse)
	}
	result := results.NewOutputResultFromAnalysisResult(analysisResult)
	r.printFormattingStart(output)
	errFinish := formatter.Finish(result, output, formatterInput)
	if errFinish != nil {
		r.printFormatterError(output, string(formatter.GetName()), errFinish)
	}
	if options.FailOnUncovered && result.HasUncovered() {
		return apperrors.NewCommandRunExceptionFinishedWithUncovered()
	}
	if result.HasViolations() {
		return apperrors.NewCommandRunExceptionFinishedWithViolations()
	}
	if result.HasErrors() {
		return apperrors.NewCommandRunExceptionFailedWithErrors()
	}

	return nil
}

func (r *AnalyseRunner) printCollectViolations(output results.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: "<info>collecting violations.</>"})
	}
}

func (r *AnalyseRunner) printFormattingStart(output results.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: "<info>formatting dependencies.</>"})
	}
}

func (r *AnalyseRunner) printFormatterError(output results.OutputInterface, formatterName string, error error) {
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
	output.GetStyle().Error(results.StringOrArrayOfStrings{Strings: []string{"", fmt.Sprintf("OutputInterface formatter %s threw an Exception:", formatterName), fmt.Sprintf("Message: %s", error.Error()), ""}})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
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

func (r *AnalyseRunner) printAnalysisException(output results.OutputInterface, exception *multierror.Error) {
	message := []string{"Analysis finished with an Exception.", JsonMultiErrFormatFunc(exception.Errors), ""}
	output.GetStyle().Error(results.StringOrArrayOfStrings{Strings: message})
}

func (r *AnalyseRunner) printFormatterNotFoundException(output results.OutputInterface, formatterName string) {
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})

	knownFormatters := make([]string, 0)
	for _, formatterType := range r.formatterProvider.GetKnownFormatters() {
		knownFormatters = append(knownFormatters, string(formatterType))
	}

	output.GetStyle().Error(results.StringOrArrayOfStrings{Strings: []string{fmt.Sprintf("Output formatter %s not found.", formatterName), "Available formatters:", strings.Join(knownFormatters, ", "), ""}})
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: ""})
}
