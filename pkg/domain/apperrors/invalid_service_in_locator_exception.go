package apperrors

import (
	"fmt"
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

func NewInvalidServiceInLocatorExceptionInvalidType(id string, actualType string, expectedTypes ...string) *InvalidServiceInLocatorException {
	message := fmt.Sprintf("Trying to get unsupported service \"%s\" from locator (expected \"%s\", but is \"%s\").", id, actualType, strings.Join(expectedTypes, "\", \""))
	return newInvalidServiceInLocatorException(message)
}
