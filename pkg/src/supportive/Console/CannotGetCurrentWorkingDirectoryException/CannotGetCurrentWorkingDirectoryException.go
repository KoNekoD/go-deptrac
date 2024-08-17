package CannotGetCurrentWorkingDirectoryException

type CannotGetCurrentWorkingDirectoryException struct {
	Message string
}

func (e *CannotGetCurrentWorkingDirectoryException) Error() string {
	return e.Message
}

func newCannotGetCurrentWorkingDirectoryException(message string) *CannotGetCurrentWorkingDirectoryException {
	return &CannotGetCurrentWorkingDirectoryException{
		Message: message,
	}
}

func NewCannotGetCurrentWorkingDirectoryExceptionCannotGetCWD() *CannotGetCurrentWorkingDirectoryException {
	return newCannotGetCurrentWorkingDirectoryException(
		"Could not get current working directory. Check `getcwd()` internal PHP function for details.",
	)
}
