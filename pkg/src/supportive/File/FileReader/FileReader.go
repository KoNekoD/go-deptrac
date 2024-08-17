package FileReader

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/Exception/CouldNotReadFileException"
	"os"
)

func FileReaderRead(fileName string) (string, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return "", CouldNotReadFileException.NewCouldNotReadFileExceptionFromFilename(fileName, err)
	}
	return string(contents), nil
}
