package assert

import (
	"fmt"
	"io"
	"reflect"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}

func reflectIsNil(v any) bool {
	defer func() {
		_ = recover() // not nil if IsNil below panics
	}()
	return reflect.ValueOf(v).IsNil()
}

func Nil(v any) {
	if v == nil {
		return
	}
	if !reflectIsNil(v) {
		panic(fmt.Sprint("not nil: ", v))
	}
}

func Ok[T any](v T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("not ok: %v, %v", v, err))
	}
	return v
}

func True(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func Truef(cond bool, format string, a ...any) {
	if !cond {
		panic(fmt.Sprintf(format, a...))
	}
}
