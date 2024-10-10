package apperrors

import "fmt"

type CannotLoadConfiguration struct {
	Message string
}

func (e *CannotLoadConfiguration) Error() string {
	return e.Message
}

func NewCannotLoadConfiguration(filename string, message string) *CannotLoadConfiguration {
	return &CannotLoadConfiguration{fmt.Sprintf("Could not load %s. Reason: %s", filename, message)}
}
