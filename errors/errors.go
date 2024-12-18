package errors

import (
	"fmt"
	"strings"
)

func Handle(err *error, msg string) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", msg, *err)
	}
}

func Handlef(err *error, format string, a ...any) {
	if *err != nil {
		*err = fmt.Errorf(format + ": %w", append(a, *err)...)
	}
}

type multiError struct {
	errs []error
}

func (merr *multiError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d errors occurred:", len(merr.errs))
	for i, err := range merr.errs {
		fmt.Fprintf(&b, "\n%d. %s", i+1, err)
	}
	return b.String()
}

func (merr *multiError) Unwrap() []error {
	return merr.errs
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
		if merr, ok := errs[0].(*multiError); ok {
			return &multiError{append(merr.errs, errs[1:i]...)}
		}
		return &multiError{errs[:i:i]}
	}
}
