package types

func Ptr[T any](t T) *T {
	return &t
}
