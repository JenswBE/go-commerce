package generics

func Must[T any](input T, err error) T {
	if err != nil {
		panic(err)
	}
	return input
}
