package ergo

import "fmt"

func Annotate(err *error, s string) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", s, *err)
	}
}

func Check0(err error) {
	if err != nil {
		panic(err)
	}
}

func Check1[T1 any](arg1 T1, err error) T1 {
	Check0(err)
	return arg1
}

func Check2[T1 any, T2 any](arg1 T1, arg2 T2, err error) (T1, T2) {
	Check0(err)
	return arg1, arg2
}

func Check3[T1 any, T2 any, T3 any](arg1 T1, arg2 T2, arg3 T3, err error) (T1, T2, T3) {
	Check0(err)
	return arg1, arg2, arg3
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
