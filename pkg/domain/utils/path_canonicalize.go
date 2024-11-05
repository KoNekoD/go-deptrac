package utils

import "path/filepath"

func PathCanonicalize(path string) string {
	result, err := filepath.EvalSymlinks(path)

	if err != nil {
		panic(err)
	}

	return result
}
