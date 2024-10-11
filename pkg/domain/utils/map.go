package utils

func MapSlice[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Map[T1 comparable, T2, V any](ts map[T1]T2, fn func(T1, T2) V) []V {
	result := make([]V, len(ts))
	i := 0
	for k, v := range ts {
		result[i] = fn(k, v)
		i++
	}
	return result
}
