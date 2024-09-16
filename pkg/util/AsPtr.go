package util

func AsPtr[T any](value T) *T {
	return &value
}
