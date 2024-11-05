package apperrors

import (
	"fmt"
)

type FileAlreadyExistsException struct {
	Message string
}

func (e *FileAlreadyExistsException) Error() string {
	return e.Message
}

func newFileAlreadyExistsException(message string) *FileAlreadyExistsException {
	return &FileAlreadyExistsException{Message: message}
}

// NewFileAlreadyExistsExceptionAlreadyExists - use via utils.PathCanonicalize(file.Name())
func NewFileAlreadyExistsExceptionAlreadyExists(path string) *FileAlreadyExistsException {
	return newFileAlreadyExistsException(fmt.Sprintf("A file_supportive named \"%s\" already exists. ", path))
}
