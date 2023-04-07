package testutil

func PtrOf[T any](i T) *T {
	return &i
}
