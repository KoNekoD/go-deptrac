package formatters

import (
	results2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/results"
)

type OutputFormatterInterface interface {
	// GetName - used as an identifier to access to the formatter or to display something more user-friendly to the user when referring to the formatter
	GetName() enums.OutputFormatterType

	// Finish - Renders the final result_contract.
	Finish(result *results2.OutputResult, output results.OutputInterface, outputFormatterInput *OutputFormatterInput) error
}
