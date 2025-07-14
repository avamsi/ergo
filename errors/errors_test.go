package errors_test

import (
	stderrors "errors"
	"slices"
	"testing"

	"github.com/avamsi/ergo/errors"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		name string
		err  error
		msg  string
		want string
	}{
		{
			name: "nil",
			err:  nil,
			msg:  "msg",
			want: "",
		},
		{
			name: "non-nil",
			err:  stderrors.New("err"),
			msg:  "msg",
			want: "msg: err",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.err // copy so we don't modify the original
			errors.Handle(&err, test.msg)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Handle(...) = %q, want %q", err, test.want)
			}
			// Handle is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := stderrors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%q) = %q, want %q", err, got, test.err)
			}
		})
	}
}

func TestHandlef(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		format string
		args   []any
		want   string
	}{
		{
			name:   "nil",
			err:    nil,
			format: "msg %d",
			args:   []any{1},
			want:   "",
		},
		{
			name:   "non-nil",
			err:    stderrors.New("err"),
			format: "msg %d",
			args:   []any{2},
			want:   "msg 2: err",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.err // copy so we don't modify the original
			errors.Handlef(&err, test.format, test.args...)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Handlef(...) = %q, want %q", err, test.want)
			}
			// Handlef is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := stderrors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%q) = %q, want %q", err, got, test.err)
			}
		})
	}
}

func TestJoinNil(t *testing.T) {
	if err := errors.Join(); err != nil {
		t.Errorf("Join() = %q, want nil", err)
	}
	if err := errors.Join(nil); err != nil {
		t.Errorf("Join(nil) = %q, want nil", err)
	}
	if err := errors.Join(nil, nil); err != nil {
		t.Errorf("errors.Join(nil, nil) = %q, want nil", err)
	}
}

func TestJoinUnwrap(t *testing.T) {
	var (
		err1  = stderrors.New("err1")
		err2  = stderrors.New("err2")
		tests = []struct {
			errs []error
			want []error
		}{
			{
				errs: []error{err1, err2},
				want: []error{err1, err2},
			},
			{
				errs: []error{err1, nil, err2},
				want: []error{err1, err2},
			},
		}
	)
	for _, test := range tests {
		got := errors.Join(test.errs...).(interface{ Unwrap() []error }).Unwrap()
		if !slices.Equal(got, test.want) {
			t.Errorf("Join(%q) = %q, want %q", test.errs, got, test.want)
		}
	}
}

func TestJoinError(t *testing.T) {
	var (
		err1  = stderrors.New("err1")
		err2  = stderrors.New("err2")
		err3  = stderrors.New("err3")
		tests = []struct {
			errs []error
			want string
		}{
			{
				errs: []error{err1},
				want: "err1",
			},
			{
				errs: []error{err1, err2},
				want: "2 errors occurred:\n   1. err1\n   2. err2",
			},
			{
				errs: []error{err1, nil, err2},
				want: "2 errors occurred:\n   1. err1\n   2. err2",
			},
			{
				errs: []error{err1, errors.Join(err2, err3)},
				want: "2 errors occurred:\n" +
					"   1. err1\n" +
					"   2. 2 errors occurred:\n   1. err2\n   2. err3",
			},
			{
				errs: []error{errors.Join(err1, err2), err3},
				want: "3 errors occurred:\n   1. err1\n   2. err2\n   3. err3",
			},
		}
	)
	for _, test := range tests {
		got := errors.Join(test.errs...).Error()
		if got != test.want {
			t.Errorf("Join(%q) = %q, want %q", test.errs, got, test.want)
		}
	}
}
