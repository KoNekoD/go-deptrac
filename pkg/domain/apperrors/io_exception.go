package apperrors

type IOException struct {
	Message string
}

func (e *IOException) Error() string {
	return e.Message
}

func newIOException(message string) *IOException {
	return &IOException{Message: message}
}

func NewIOExceptionCouldNotCopy(message string) *IOException {
	return newIOException(message)
}
