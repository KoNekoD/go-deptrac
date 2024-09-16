package InvalidServiceInLocatorException

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	"strings"
)

type InvalidServiceInLocatorException struct {
	Message string
}

func (e *InvalidServiceInLocatorException) Error() string {
	return e.Message
}

func newInvalidServiceInLocatorException(message string) *InvalidServiceInLocatorException {
	return &InvalidServiceInLocatorException{Message: message}
}

func NewInvalidServiceInLocatorExceptionInvalidType(id output_formatter_contract.OutputFormatterType, actualType string, expectedTypes ...string) *InvalidServiceInLocatorException {
	message := fmt.Sprintf("Trying to get unsupported service \"%s\" from locator (expected \"%s\", but is \"%s\").", id, actualType, strings.Join(expectedTypes, "\", \""))
	return newInvalidServiceInLocatorException(message)
}
