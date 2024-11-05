package utils

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"os"
)

func FileReaderRead(fileName string) (string, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return "", apperrors.NewCouldNotReadFileExceptionFromFilename(fileName, err)
	}
	return string(contents), nil
}
