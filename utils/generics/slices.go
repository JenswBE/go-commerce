package generics

func Map[T any](slice []T, mapFunc func(T) T) []T {
	for i, value := range slice {
		slice[i] = mapFunc(value)
	}
	return slice
}

func Last[T any](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}
	return slice[len(slice)-1]
}
