package util

func MapKeyExists[Key comparable, Value any](v map[Key]Value, k Key) bool {
	_, ok := v[k]
	return ok
}
