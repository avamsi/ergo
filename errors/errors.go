package errors

import (
	"fmt"
	"strings"
)

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

type joinError struct {
	errs []error
}

func (err *joinError) Error() string {
	var b strings.Builder
	for _, err := range err.errs {
		b.WriteString("\n\t")
		b.WriteString(err.Error())
	}
	return b.String()
}

func (err *joinError) Unwrap() []error {
	return err.errs
}

func Join(errs ...error) error {
	var i int
	for _, err := range errs {
		if err != nil {
			errs[i] = err
			i++
		}
	}
	switch i {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return &joinError{errs[:i:i]}
	}
}
