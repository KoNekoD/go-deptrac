package exception

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
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

func NewFileAlreadyExistsExceptionAlreadyExists(file *os.File) *FileAlreadyExistsException {
	return newFileAlreadyExistsException(fmt.Sprintf("A file_supportive named \"%s\" already exists. ", util.PathCanonicalize(file.Name())))
}
