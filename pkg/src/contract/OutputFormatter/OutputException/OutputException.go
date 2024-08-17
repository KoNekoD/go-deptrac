package OutputException

// OutputException - Thrown when you are unable to provide output with your custom OutputFormatter.
type OutputException struct {
	Message string
}

func NewOutputExceptionWithMessage(message string) *OutputException {
	return &OutputException{Message: message}
}
