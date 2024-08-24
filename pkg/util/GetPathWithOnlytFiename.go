package util

import (
	"errors"
	"os"
	"path/filepath"
)

func GetPathWithOnlyFilename(path string) (string, error) {
	pathStat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if pathStat.IsDir() {
		return "", errors.New("required path is a file")
	}

	path = filepath.Base(path)

	return path, nil
}
