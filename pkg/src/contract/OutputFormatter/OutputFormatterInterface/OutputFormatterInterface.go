package OutputFormatterInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInput"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/OutputResult"
)

type OutputFormatterInterface interface {
	// GetName - used as an identifier to access to the formatter or to display something more user-friendly to the user when referring to the formatter
	GetName() OutputFormatterType.OutputFormatterType

	// Finish - Renders the final result.
	Finish(result *OutputResult.OutputResult, output OutputInterface.OutputInterface, outputFormatterInput *OutputFormatterInput.OutputFormatterInput) error
}
