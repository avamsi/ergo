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

func Check2[T1 any, T2 any](arg1 T1, arg2 T2, err error) (T1, T2) {
	Check0(err)
	return arg1, arg2
}

func Check3[T1 any, T2 any, T3 any](arg1 T1, arg2 T2, arg3 T3, err error) (T1, T2, T3) {
	Check0(err)
	return arg1, arg2, arg3
}
