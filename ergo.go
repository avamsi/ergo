package ergo

import "fmt"

func Must0(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func Must1[T1 any](arg1 T1, err error) T1 {
	if err == nil {
		return arg1
	}
	panic(err)
}

func Must2[T1 any, T2 any](arg1 T1, arg2 T2, err error) (T1, T2) {
	if err == nil {
		return arg1, arg2
	}
	panic(err)
}

func Must3[T1 any, T2 any, T3 any](arg1 T1, arg2 T2, arg3 T3, err error) (T1, T2, T3) {
	if err == nil {
		return arg1, arg2, arg3
	}
	panic(err)
}

func Error1[T1 any](_ T1, err error) error {
	return err
}

func Error2[T1 any, T2 any](_ T1, _ T2, err error) error {
	return err
}

func Error3[T1 any, T2 any, T3 any](_ T1, _ T2, _ T3, err error) error {
	return err
}
