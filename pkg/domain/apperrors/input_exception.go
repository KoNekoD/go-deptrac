package apperrors

import "fmt"

type InputException struct {
	Message  string
	Previous error
}

func (i *InputException) Error() string {
	if i.Previous != nil {
		return fmt.Sprintf("%s\n%s", i.Message, i.Previous.Error())
	} else {
		return i.Message
	}
}

func NewInputExceptionCouldNotCollectFiles(previous error) *InputException {
	return &InputException{Message: "Could not collect files.", Previous: previous}
}
