package assert_test

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/avamsi/ergo/assert"
)

type errCloser struct{}

func (c errCloser) Close() error {
	return errors.New("err")
}

type nopCloser struct{}

func (c nopCloser) Close() error {
	return nil
}

func TestPanic(t *testing.T) {
	tests := []struct {
		name string
		f    func()
		want string
	}{
		{
			name: "not-closed",
			f: func() {
				assert.Close(errCloser{})
			},
			want: "err",
		},
		{
			name: "int-not-nil",
			f: func() {
				assert.Nil(42)
			},
			want: "not nil: 42",
		},
		{
			name: "string-not-nil",
			f: func() {
				assert.Nil("boo")
			},
			want: "not nil: boo",
		},
		{
			name: "error-not-nil",
			f: func() {
				assert.Nil(errors.New("err"))
			},
			want: "not nil: err",
		},
		{
			name: "map-not-nil",
			f: func() {
				assert.Nil(map[int]int{42: 69})
			},
			want: "not nil: map[42:69]",
		},
		{
			name: "slice-not-nil",
			f: func() {
				assert.Nil([]int{42, 69})
			},
			want: "not nil: [42 69]",
		},
		{
			name: "struct-pointer-not-nil",
			f: func() {
				assert.Nil(&struct{ x, y int }{42, 69})
			},
			want: "not nil: &{42 69}",
		},
		{
			name: "int-error-not-ok",
			f: func() {
				assert.Ok(42, errors.New("err"))
			},
			want: "not ok: 42, err",
		},
		{
			name: "slice-error-not-ok",
			f: func() {
				assert.Ok([]int{42, 69}, errors.New("err"))
			},
			want: "not ok: [42 69], err",
		},
		{
			name: "struct-error-not-ok",
			f: func() {
				assert.Ok(struct{ x, y int }{42, 69}, errors.New("err"))
			},
			want: "not ok: {42 69}, err",
		},
		{
			name: "false-not-true",
			f: func() {
				assert.True(false, "false")
			},
			want: "false",
		},
		{
			name: "false-not-true-f",
			f: func() {
				assert.Truef(false, "false-%s", "f")
			},
			want: "false-f",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				got := recover()
				if err, ok := got.(error); ok {
					got = err.Error()
				}
				if got != test.want {
					t.Errorf("got panic %#v, want %#v", got, test.want)
				}
			}()
			test.f()
		})
	}
}

func TestNotPanic(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "closed",
			f: func() {
				assert.Close(nopCloser{})
			},
		},
		{
			name: "nil",
			f: func() {
				assert.Nil(nil)
			},
		},
		{
			name: "zero-error-nil",
			f: func() {
				var err error
				assert.Nil(err)
			},
		},
		{
			name: "zero-map-nil",
			f: func() {
				var m map[int]int
				assert.Nil(m)
			},
		},
		{
			name: "zero-slice-nil",
			f: func() {
				var s []int
				assert.Nil(s)
			},
		},
		{
			name: "zero-pointer-nil",
			f: func() {
				var p *int
				assert.Nil(p)
			},
		},
		{
			name: "error-nil-ok",
			f: func() {
				assert.Ok(42, nil)
			},
		},
		{
			name: "true",
			f: func() {
				assert.True(true, "true")
			},
		},
		{
			name: "true-f",
			f: func() {
				assert.Truef(true, "true-%s", "f")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if got := recover(); got != nil {
					t.Errorf("want no panic, got: %#v\n%s", got, debug.Stack())
				}
			}()
			test.f()
		})
	}
}
