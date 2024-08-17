package util

import "os"

func IsReadable(filepath string) bool {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	permMode := fileInfo.Mode().Perm()

	return permMode&0444 == 0444
}
