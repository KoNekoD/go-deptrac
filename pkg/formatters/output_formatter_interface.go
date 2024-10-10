package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/results"
)

type OutputFormatterInterface interface {
	// GetName - used as an identifier to access to the formatter or to display something more user-friendly to the user when referring to the formatter
	GetName() OutputFormatterType

	// Finish - Renders the final result_contract.
	Finish(result *results.OutputResult, output results.OutputInterface, outputFormatterInput *OutputFormatterInput) error
}
