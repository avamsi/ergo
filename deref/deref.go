package deref

func Or[T any](ptr *T, value T) T {
	if ptr != nil {
		return *ptr
	}
	return value
}
