package config

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseOptions"
)

type analyseOptionsHook struct{}
type AnalyseOptionsHook interface {
	GetOptions() *AnalyseOptions.AnalyseOptions
}

func NewAnalyseOptionsHook() AnalyseOptionsHook {
	return &analyseOptionsHook{}
}

// OptionFormatterUsage - TODO: Add Possible: ["%s"]', \implode('", "', $this->formatterProvider->getKnownFormatters()))
const (
	OptionReportUncovered      = "report-uncovered"
	OptionReportUncoveredUsage = "Report uncovered dependencies"
	OptionFailOnUncovered      = "fail-on-uncovered"
	OptionFailOnUncoveredUsage = "Fails if any uncovered dependency is found"
	OptionReportSkipped        = "report-skipped"
	OptionReportSkippedUsage   = "Report skipped violations"
	OptionFormatter            = "formatter"
	OptionFormatterUsage       = "Format in which to print the result of the analysis"
	OptionOutput               = "output"
	OptionOutputUsage          = "Output file path for formatter (if applicable)"
	OptionNoProgress           = "no-progress"
	OptionNoProgressUsage      = "Do not show progress bar"
)

func (h *analyseOptionsHook) GetOptions() *AnalyseOptions.AnalyseOptions {

	reportUncovered := flag.Bool(OptionReportUncovered, false, OptionReportUncoveredUsage)
	failOnUncovered := flag.Bool(OptionFailOnUncovered, false, OptionFailOnUncoveredUsage)
	reportSkipped := flag.Bool(OptionReportSkipped, false, OptionReportSkippedUsage)
	formatter := flag.String(OptionFormatter, "", OptionFormatterUsage)
	output := flag.String(OptionOutput, "", OptionOutputUsage)
	noProgress := flag.Bool(OptionNoProgress, false, OptionNoProgressUsage)

	return &AnalyseOptions.AnalyseOptions{
		ReportUncovered: *reportUncovered,
		FailOnUncovered: *failOnUncovered,
		ReportSkipped:   *reportSkipped,
		Formatter:       *formatter,
		Output:          output,
		NoProgress:      *noProgress,
	}
}
