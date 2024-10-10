package options

type AnalyseOptions struct {
	NoProgress      bool
	Formatter       string
	Output          *string
	ReportSkipped   bool
	ReportUncovered bool
	FailOnUncovered bool
}

func NewAnalyseOptions(noProgress bool, formatter string, output *string, reportSkipped bool, reportUncovered bool, failOnUncovered bool) *AnalyseOptions {
	return &AnalyseOptions{
		NoProgress:      noProgress,
		Formatter:       formatter,
		Output:          output,
		ReportSkipped:   reportSkipped,
		ReportUncovered: reportUncovered,
		FailOnUncovered: failOnUncovered,
	}
}
