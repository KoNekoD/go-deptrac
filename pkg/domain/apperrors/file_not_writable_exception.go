package apperrors

import (
	"fmt"
)

type FileNotWritableException struct {
	Message string
}

func (e *FileNotWritableException) Error() string {
	return e.Message
}

func newFileNotWritableException(message string) *FileNotWritableException {
	return &FileNotWritableException{Message: message}
}

// NewFileNotWritableExceptionFromFilePath - use via PathCanonicalize(file.Name())
func NewFileNotWritableExceptionFromFilePath(path string) *FileNotWritableException {
	return newFileNotWritableException(fmt.Sprintf("Could not write file_supportive \"%s\".", path))
}
