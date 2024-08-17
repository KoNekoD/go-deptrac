package util

func MapKeyIsArrayOfStrings[Key comparable](v map[Key]interface{}, k Key) bool {
	mapKeyValue := v[k]

	values, ok1 := mapKeyValue.([]interface{})

	if !ok1 {
		return false
	}

	for _, value := range values {
		_, ok2 := value.(string)

		if !ok2 {
			return false
		}
	}

	return true
}
