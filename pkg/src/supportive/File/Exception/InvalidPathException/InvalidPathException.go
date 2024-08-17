package InvalidPathException

import (
	"fmt"
	"os"
)

type InvalidPathException struct {
	Message string
}

func (e *InvalidPathException) Error() string {
	return e.Message
}

func newInvalidPathException(message string) *InvalidPathException {
	return &InvalidPathException{Message: message}
}

func NewInvalidPathExceptionUnreadablePath(path os.FileInfo) *InvalidPathException {
	return newInvalidPathException(fmt.Sprintf("Path \"%s\" is not a directory or is not readable.", path.Name()))
}
