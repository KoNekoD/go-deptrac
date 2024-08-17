package OutputInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
)

// OutputInterface - Wrapper around Symfony OutputInterface.
type OutputInterface interface {
	WriteFormatted(message string)
	WriteLineFormatted(message OutputStyleInterface.StringOrArrayOfStrings)
	WriteRaw(message string)
	GetStyle() OutputStyleInterface.OutputStyleInterface
	IsVerbose() bool
	IsDebug() bool
}
