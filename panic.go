package ergo

import "fmt"

func Panic(a ...any) {
	panic(fmt.Sprint(a...))
}

func Panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
