package check

import (
	"fmt"
	"reflect"
)

func isNil(v any) bool {
	if v == nil {
		return true
	}
	defer func() {
		recover() // not nil if IsNil below panics
	}()
	return reflect.ValueOf(v).IsNil()
}

func Nil(v any) {
	if !isNil(v) {
		panic(fmt.Sprint("not nil: ", v))
	}
}

func Ok[T any](arg T, err error) T {
	if err != nil {
		panic(err)
	}
	return arg
}

func True(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func Truef(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(format, args...))
	}
}
