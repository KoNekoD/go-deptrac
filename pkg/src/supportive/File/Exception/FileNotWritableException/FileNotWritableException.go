package FileNotWritableException

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
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

func NewFileNotWritableExceptionFromFilePath(file *os.File) *FileNotWritableException {
	return newFileNotWritableException(fmt.Sprintf("Could not write file \"%s\".", util.PathCanonicalize(file.Name())))
}
