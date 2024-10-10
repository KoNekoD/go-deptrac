package apperrors

import "fmt"

type FileNotExistsException struct {
	Message string
}

func (e *FileNotExistsException) Error() string {
	return e.Message
}

func newFileNotExistsException(message string) *FileNotExistsException {
	return &FileNotExistsException{Message: message}
}

func NewFileNotExistsExceptionFromFilePath(filepath string) *FileNotExistsException {
	return newFileNotExistsException(fmt.Sprintf("\"%s\" is not a valid path or does not exists.", filepath))
}
