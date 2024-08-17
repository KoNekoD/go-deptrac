package ExceptionInterface

// ExceptionInterface - Shared interface for all Exceptions that Deptrac can possibly throw. You can use this to ensure that no exceptions go unhandled when integrating with Deptrac codebase.
type ExceptionInterface interface{}

type Exception struct {
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

func NewException(message string) *Exception {
	return &Exception{
		Message: message,
	}
}
