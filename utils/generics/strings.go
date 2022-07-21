package generics

func Map[T any](slice []T, mapFunc func(T) T) []T {
	for i, value := range slice {
		slice[i] = mapFunc(value)
	}
	return slice
}
