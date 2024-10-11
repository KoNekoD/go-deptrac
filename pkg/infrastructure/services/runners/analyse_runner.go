package runners

import (
	"encoding/json"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/options"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
	"github.com/hashicorp/go-multierror"
	"strings"
)

// AnalyseRunner - Should only be used by AnalyseCommand
type AnalyseRunner struct {
	analyzer          *analysers.DependencyLayersAnalyser
	formatterProvider *formatters.FormatterProvider
}

func NewAnalyseRunner(analyzer *analysers.DependencyLayersAnalyser, formatterProvider *formatters.FormatterProvider) *AnalyseRunner {
	return &AnalyseRunner{
		analyzer:          analyzer,
		formatterProvider: formatterProvider,
	}
}

func (r *AnalyseRunner) Run(options *options.AnalyseOptions, output services.OutputInterface) error {
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

func (r *AnalyseRunner) printCollectViolations(output services.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(services.StringOrArrayOfStrings{String: "<info>collecting violations.</>"})
	}
}

func (r *AnalyseRunner) printFormattingStart(output services.OutputInterface) {
	if output.IsVerbose() {
		output.WriteLineFormatted(services.StringOrArrayOfStrings{String: "<info>formatting dependencies.</>"})
	}
}

func (r *AnalyseRunner) printFormatterError(output services.OutputInterface, formatterName string, error error) {
	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: ""})
	output.GetStyle().Error(services.StringOrArrayOfStrings{Strings: []string{"", fmt.Sprintf("OutputInterface formatter %s threw an Exception:", formatterName), fmt.Sprintf("Message: %s", error.Error()), ""}})
	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: ""})
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

func (r *AnalyseRunner) printAnalysisException(output services.OutputInterface, exception *multierror.Error) {
	message := []string{"Analysis finished with an Exception.", JsonMultiErrFormatFunc(exception.Errors), ""}
	output.GetStyle().Error(services.StringOrArrayOfStrings{Strings: message})
}

func (r *AnalyseRunner) printFormatterNotFoundException(output services.OutputInterface, formatterName string) {
	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: ""})

	knownFormatters := make([]string, 0)
	for _, formatterType := range r.formatterProvider.GetKnownFormatters() {
		knownFormatters = append(knownFormatters, string(formatterType))
	}

	output.GetStyle().Error(services.StringOrArrayOfStrings{Strings: []string{fmt.Sprintf("Output formatter %s not found.", formatterName), "Available formatters:", strings.Join(knownFormatters, ", "), ""}})
	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: ""})
}
