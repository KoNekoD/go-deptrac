package output_formatter_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/result_contract/output_result"
)

type OutputFormatterInterface interface {
	// GetName - used as an identifier to access to the formatter or to display something more user-friendly to the user when referring to the formatter
	GetName() OutputFormatterType

	// Finish - Renders the final result_contract.
	Finish(result *output_result.OutputResult, output OutputInterface, outputFormatterInput *OutputFormatterInput) error
}
