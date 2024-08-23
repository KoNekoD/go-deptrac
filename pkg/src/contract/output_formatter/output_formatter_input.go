package output_formatter

type OutputFormatterInput struct {
	OutputPath      *string
	ReportSkipped   bool
	ReportUncovered bool
	FailOnUncovered bool
}

func NewOutputFormatterInput(outputPath string, reportSkipped bool, reportUncovered bool, failOnUncovered bool) *OutputFormatterInput {
	return &OutputFormatterInput{
		OutputPath:      &outputPath,
		ReportSkipped:   reportSkipped,
		ReportUncovered: reportUncovered,
		FailOnUncovered: failOnUncovered,
	}
}
