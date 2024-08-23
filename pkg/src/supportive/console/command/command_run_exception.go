package command

import "fmt"

type CommandRunException struct {
	Message  string
	Previous error
}

func (e *CommandRunException) Error() string {
	if e.Previous != nil {
		return fmt.Sprintf("%s%s", e.Message, e.Previous.Error())
	}
	return e.Message
}

func newCommandRunException(message string, previous error) *CommandRunException {
	return &CommandRunException{Message: message, Previous: previous}
}

func NewCommandRunExceptionInvalidFormatter() *CommandRunException {
	return newCommandRunException("Invalid output formatter selected.", nil)
}

func NewCommandRunExceptionFinishedWithUncovered() *CommandRunException {
	return newCommandRunException("Analysis finished, but contains uncovered tokens.", nil)
}

func NewCommandRunExceptionFinishedWithViolations() *CommandRunException {
	return newCommandRunException("Analysis finished, but contains ruleset violations.", nil)
}

func NewCommandRunExceptionFailedWithErrors() *CommandRunException {
	return newCommandRunException("Analysis failed, due to an error.", nil)
}

func NewCommandRunExceptionAnalyserException(analyserException error) *CommandRunException {
	return newCommandRunException("Analysis failed.", analyserException)
}
