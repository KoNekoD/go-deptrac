package apperrors

import "fmt"

type CouldNotReadFileException struct {
	Message  string
	Previous error
}

func (e *CouldNotReadFileException) Error() string {
	if e.Previous != nil {
		return fmt.Sprintf("%s%s", e.Message, e.Previous.Error())
	}
	return e.Message
}

func newCouldNotReadFileException(message string, previous error) *CouldNotReadFileException {
	return &CouldNotReadFileException{Message: message, Previous: previous}
}

func NewCouldNotReadFileExceptionFromFilename(filename string, previous error) *CouldNotReadFileException {
	return newCouldNotReadFileException(fmt.Sprintf("File \"%s\" cannot be read. ", filename), previous)
}
