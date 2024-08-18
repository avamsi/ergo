package errors_test

import (
	stderrors "errors"
	"testing"

	"github.com/avamsi/ergo/errors"
)

func TestAnnotate(t *testing.T) {
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
			errors.Annotate(&err, test.msg)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Annotate(...) = %q, want %q", err, test.want)
			}
			// Annotate is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := stderrors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%q) = %q, want %q", err, got, test.err)
			}
		})
	}
}

func TestAnnotatef(t *testing.T) {
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
			errors.Annotatef(&err, test.format, test.args...)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Annotatef(...) = %q, want %q", err, test.want)
			}
			// Annotatef is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := stderrors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%q) = %q, want %q", err, got, test.err)
			}
		})
	}
}
