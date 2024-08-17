package util

import "slices"

func IsPrimitiveType(v string) bool {
	primitiveTypes := []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"bool",
		"string",
	}

	return slices.Contains(primitiveTypes, v)
}
