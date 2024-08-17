package util

func MapKeyIsInt[Key comparable](v map[Key]interface{}, k Key) bool {
	mapKeyValue := v[k]

	_, ok := mapKeyValue.(int)

	return ok
}
