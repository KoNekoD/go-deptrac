package file

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/file/exception"
	"os"
)

func FileReaderRead(fileName string) (string, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return "", exception.NewCouldNotReadFileExceptionFromFilename(fileName, err)
	}
	return string(contents), nil
}
