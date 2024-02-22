package repositories

func ToEntitiesList[T, R any](input []T, convFunc func(T) R) []R {
	output := make([]R, len(input))
	for i, item := range input {
		output[i] = convFunc(item)
	}
	return output
}
