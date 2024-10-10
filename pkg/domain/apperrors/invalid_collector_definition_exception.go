package apperrors

import (
	"fmt"
	"strings"
)

// InvalidCollectorDefinitionException - Thrown when the configuration of a particular collector is not valid. Use this error when writing custom collectors.
type InvalidCollectorDefinitionException struct {
	Message  string
	Previous error
}

func (e InvalidCollectorDefinitionException) Error() string {
	if e.Previous != nil {
		return fmt.Sprintf("%s\n%s", e.Message, e.Previous.Error())
	} else {
		return e.Message
	}
}

func NewInvalidCollectorDefinitionMissingType() *InvalidCollectorDefinitionException {
	msg := "Could not resolve collector definition because of missing \"type\" field."

	return &InvalidCollectorDefinitionException{Message: msg}
}

// NewInvalidCollectorDefinitionUnsupportedType - use via types CollectorType, supportedTypes []CollectorType, previous error
func NewInvalidCollectorDefinitionUnsupportedType(needed string, supportedTypes []string, previous error) *InvalidCollectorDefinitionException {
	possible := strings.Join(supportedTypes, "\", \"")

	msg := fmt.Sprintf("Could not find a collector for type \"%s\". Supported types: \"%s\".", needed, possible)

	return &InvalidCollectorDefinitionException{Message: msg, Previous: previous}
}

func NewInvalidCollectorDefinitionInvalidCollectorConfiguration(message string) *InvalidCollectorDefinitionException {
	return &InvalidCollectorDefinitionException{Message: message}
}
