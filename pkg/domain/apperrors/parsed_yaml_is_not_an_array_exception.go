package apperrors

import "fmt"

type ParsedYamlIsNotAnArrayException struct {
	Message string
}

func (e *ParsedYamlIsNotAnArrayException) Error() string {
	return e.Message
}

func newParsedYamlIsNotAnArrayException(message string) *ParsedYamlIsNotAnArrayException {
	return &ParsedYamlIsNotAnArrayException{Message: message}
}

func NewParsedYamlIsNotAnArrayExceptionFromFilename(filename string) *ParsedYamlIsNotAnArrayException {
	return newParsedYamlIsNotAnArrayException(fmt.Sprintf("File \"%s\" can be parsed as YAML, but the result_contract is not an array.", filename))
}
