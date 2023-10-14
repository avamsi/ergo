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
		fn   func()
		want string
	}{
		{
			name: "not-closed",
			fn: func() {
				assert.Close(errCloser{})
			},
			want: "err",
		},
		{
			name: "int-not-nil",
			fn: func() {
				assert.Nil(42)
			},
			want: "not nil: 42",
		},
		{
			name: "string-not-nil",
			fn: func() {
				assert.Nil("boo")
			},
			want: "not nil: boo",
		},
		{
			name: "error-not-nil",
			fn: func() {
				assert.Nil(errors.New("err"))
			},
			want: "not nil: err",
		},
		{
			name: "map-not-nil",
			fn: func() {
				assert.Nil(map[int]int{42: 69})
			},
			want: "not nil: map[42:69]",
		},
		{
			name: "slice-not-nil",
			fn: func() {
				assert.Nil([]int{42, 69})
			},
			want: "not nil: [42 69]",
		},
		{
			name: "struct-pointer-not-nil",
			fn: func() {
				assert.Nil(&struct{ x, y int }{42, 69})
			},
			want: "not nil: &{42 69}",
		},
		{
			name: "int-error-not-ok",
			fn: func() {
				assert.Ok(42, errors.New("err"))
			},
			want: "not ok: 42, err",
		},
		{
			name: "slice-error-not-ok",
			fn: func() {
				assert.Ok([]int{42, 69}, errors.New("err"))
			},
			want: "not ok: [42 69], err",
		},
		{
			name: "struct-error-not-ok",
			fn: func() {
				assert.Ok(struct{ x, y int }{42, 69}, errors.New("err"))
			},
			want: "not ok: {42 69}, err",
		},
		{
			name: "false-not-true",
			fn: func() {
				assert.True(false, "false")
			},
			want: "false",
		},
		{
			name: "false-not-true-f",
			fn: func() {
				assert.Truef(false, "false-%s", "f")
			},
			want: "false-f",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if s, ok := r.(error); ok {
					r = s.Error()
				}
				if r != test.want {
					t.Errorf("want panic %#v, got %#v\n", test.want, r)
				}
			}()
			test.fn()
		})
	}
}

func TestNotPanic(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{
			name: "closed",
			fn: func() {
				assert.Close(nopCloser{})
			},
		},
		{
			name: "nil",
			fn: func() {
				assert.Nil(nil)
			},
		},
		{
			name: "zero-error-nil",
			fn: func() {
				var err error
				assert.Nil(err)
			},
		},
		{
			name: "zero-map-nil",
			fn: func() {
				var m map[int]int
				assert.Nil(m)
			},
		},
		{
			name: "zero-slice-nil",
			fn: func() {
				var s []int
				assert.Nil(s)
			},
		},
		{
			name: "zero-pointer-nil",
			fn: func() {
				var p *int
				assert.Nil(p)
			},
		},
		{
			name: "error-nil-ok",
			fn: func() {
				assert.Ok(42, nil)
			},
		},
		{
			name: "true",
			fn: func() {
				assert.True(true, "true")
			},
		},
		{
			name: "true-f",
			fn: func() {
				assert.Truef(true, "true-%s", "f")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("want no panic, got: %#v\n%s", r, debug.Stack())
				}
			}()
			test.fn()
		})
	}
}
