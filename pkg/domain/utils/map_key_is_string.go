package utils

func MapKeyIsString[Key comparable](v map[Key]interface{}, k Key) bool {
	mapKeyValue := v[k]

	_, ok := mapKeyValue.(string)

	return ok
}
