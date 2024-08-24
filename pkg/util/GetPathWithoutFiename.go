package util

import (
	"os"
	"path/filepath"
)

func GetPathWithoutFilename(path string) (string, error) {
	pathStat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if !pathStat.IsDir() {
		path = filepath.Dir(path)
	}

	return path, nil
}
