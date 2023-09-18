package deref

func Or[T any](ptr *T, v T) T {
	if ptr != nil {
		return *ptr
	}
	return v
}
