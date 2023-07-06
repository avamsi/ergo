package errors

import "fmt"

func Annotate(err *error, msg string) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", msg, *err)
	}
}

func Annotatef(err *error, format string, args ...any) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *err)
	}
}
