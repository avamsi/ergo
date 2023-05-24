package panic

import "fmt"

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func Assertf(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(format, args...))
	}
}

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T1 any](arg1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return arg1
}

func Must2[T1 any, T2 any](arg1 T1, arg2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return arg1, arg2
}

func Must3[T1 any, T2 any, T3 any](arg1 T1, arg2 T2, arg3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return arg1, arg2, arg3
}

func Panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}
