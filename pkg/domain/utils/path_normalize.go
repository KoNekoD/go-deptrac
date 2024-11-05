package utils

import "strings"

func PathNormalize(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
