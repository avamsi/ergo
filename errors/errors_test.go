package errors

import (
	"errors"
	"testing"
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
			err:  errors.New("err"),
			msg:  "msg",
			want: "msg: err",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.err // copy so we don't modify the original
			Annotate(&err, test.msg)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Annotate(...) = %#v, want %#v\n", err.Error(), test.want)
			}
			// Annotate is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := errors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%#v) = %#v, want %#v\n", err, got, test.err)
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
			err:    errors.New("err"),
			format: "msg %d",
			args:   []any{2},
			want:   "msg 2: err",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.err // copy so we don't modify the original
			Annotatef(&err, test.format, test.args...)
			if test.err != nil && err.Error() != test.want {
				t.Errorf("Annotatef(...) = %#v, want %#v\n", err.Error(), test.want)
			}
			// Annotatef is expected to wrap the input error exactly once, so
			// unwrapping it is expected to return the original error.
			if got := errors.Unwrap(err); got != test.err {
				t.Errorf("Unwrap(%#v) = %#v, want %#v\n", err, got, test.err)
			}
		})
	}
}

func TestErrorN(t *testing.T) {
	a, b, c, err := 1, 2, errors.New("3"), errors.New("err")
	if got := Error1(a, err); got != err {
		t.Errorf("Error1(%#v, %#v) = %#v, want %#v\n", a, err, got, err)
	}
	if got := Error2(a, b, err); got != err {
		t.Errorf("Error2(%#v, %#v, %#v) = %#v, want %#v\n", a, b, err, got, err)
	}
	if got := Error3(a, b, c, err); got != err {
		t.Errorf("Error3(%#v, %#v, %#v, %#v) = %#v, want %#v\n", a, b, c, err, got, err)
	}
}
