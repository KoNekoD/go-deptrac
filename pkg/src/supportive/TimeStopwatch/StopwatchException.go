package TimeStopwatch

import "fmt"

type StopwatchException struct {
	Message string
}

func (s *StopwatchException) Error() string {
	return s.Message
}

func newStopwatchException(message string) *StopwatchException {
	return &StopwatchException{
		Message: message,
	}
}

func NewStopwatchExceptionPeriodAlreadyStarted(period string) *StopwatchException {
	return newStopwatchException(fmt.Sprintf("Period \"%s\" is already started", period))
}

func NewStopwatchExceptionPeriodNotStarted(period string) *StopwatchException {
	return newStopwatchException(fmt.Sprintf("Period \"%s\" is not started", period))
}
