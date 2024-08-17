package FileCannotBeParsedAsYamlException

import "fmt"

type FileCannotBeParsedAsYamlException struct {
	Message string
}

func (e *FileCannotBeParsedAsYamlException) Error() string {
	return e.Message
}

func newFileCannotBeParsedAsYamlException(message string) *FileCannotBeParsedAsYamlException {
	return &FileCannotBeParsedAsYamlException{Message: message}
}

func NewFileCannotBeParsedAsYamlExceptionFromFilenameAndException(filename string, exception error) *FileCannotBeParsedAsYamlException {
	return newFileCannotBeParsedAsYamlException(fmt.Sprintf("File \"%s\" cannot be parsed as YAML: %s", filename, exception.Error()))
}
