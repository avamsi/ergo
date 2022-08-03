package checks

func Check0(err error) {
	if err != nil {
		panic(err)
	}
}

func Check1[T any](arg T, err error) T {
	Check0(err)
	return arg
}
