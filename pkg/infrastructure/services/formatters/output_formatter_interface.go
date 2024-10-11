package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type OutputFormatterInterface interface {
	// GetName - used as an identifier to access to the formatter or to display something more user-friendly to the user when referring to the formatter
	GetName() enums.OutputFormatterType

	// Finish - Renders the final result_contract.
	Finish(result *results.OutputResult, output services.OutputInterface, outputFormatterInput *OutputFormatterInput) error
}
