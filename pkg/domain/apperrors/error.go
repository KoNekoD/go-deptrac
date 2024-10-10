package apperrors

type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}

func (e *Error) ToString() string {
	return e.Message
}
