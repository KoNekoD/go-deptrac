package apperrors

import (
	"fmt"
)

type UnrecognizedTokenException struct {
	Message string
}

func NewUnrecognizedTokenException(message string) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: message}
}

func (e UnrecognizedTokenException) Error() string {
	return e.Message
}

// NewCannotCreateReference - use via token TokenInterface - sprintf '%T'
func (e UnrecognizedTokenException) NewCannotCreateReference(tokenType string) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: fmt.Sprintf("Cannot create TokenReference for token '%s'", tokenType)}
}
