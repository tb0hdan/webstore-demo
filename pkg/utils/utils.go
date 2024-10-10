package utils

func Index[T any](arr []T, f func(T) bool) int {
	for i, item := range arr {
		if f(item) {
			return i
		}
	}

	return -1
}
