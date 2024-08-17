package InvalidServiceInLocatorException

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
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

func NewInvalidServiceInLocatorExceptionInvalidType(id OutputFormatterType.OutputFormatterType, actualType string, expectedTypes ...string) *InvalidServiceInLocatorException {
	message := fmt.Sprintf("Trying to get unsupported service \"%s\" from locator (expected \"%s\", but is \"%s\").", id, actualType, strings.Join(expectedTypes, "\", \""))
	return newInvalidServiceInLocatorException(message)
}
