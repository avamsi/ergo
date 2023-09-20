package errors

import "fmt"

func Annotate(err *error, msg string) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", msg, *err)
	}
}

func Annotatef(err *error, format string, a ...any) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", fmt.Sprintf(format, a...), *err)
	}
}
