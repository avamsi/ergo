package assert

import (
	"io"
	"reflect"

	"github.com/avamsi/ergo"
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
		ergo.Panic("not nil: ", v)
	}
}

func Ok[T any](v T, err error) T {
	if err != nil {
		ergo.Panicf("not ok: %v, %v", v, err)
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
		ergo.Panicf(format, a...)
	}
}
