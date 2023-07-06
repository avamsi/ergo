package check

import (
	"errors"
	"runtime/debug"
	"testing"
)

func TestPanic(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
		want string
	}{
		{
			name: "int-not-nil",
			fn: func() {
				Nil(42)
			},
			want: "not nil: 42",
		},
		{
			name: "string-not-nil",
			fn: func() {
				Nil("boo")
			},
			want: "not nil: boo",
		},
		{
			name: "error-not-nil",
			fn: func() {
				Nil(errors.New("err"))
			},
			want: "not nil: err",
		},
		{
			name: "map-not-nil",
			fn: func() {
				Nil(map[int]int{})
			},
			want: "not nil: map[]",
		},
		{
			name: "slice-not-nil",
			fn: func() {
				Nil([]int{})
			},
			want: "not nil: []",
		},
		{
			name: "error-not-ok",
			fn: func() {
				Ok(42, errors.New("err"))
			},
			want: "err",
		},
		{
			name: "false-not-true",
			fn: func() {
				True(false, "false")
			},
			want: "false",
		},
		{
			name: "false-not-true-f",
			fn: func() {
				Truef(false, "false-%s", "f")
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
			name: "nil",
			fn: func() {
				Nil(nil)
			},
		},
		{
			name: "zero-error-nil",
			fn: func() {
				var err error
				Nil(err)
			},
		},
		{
			name: "zero-map-nil",
			fn: func() {
				var m map[int]int
				Nil(m)
			},
		},
		{
			name: "zero-slice-nil",
			fn: func() {
				var s []int
				Nil(s)
			},
		},
		{
			name: "zero-pointer-nil",
			fn: func() {
				var p *int
				Nil(p)
			},
		},
		{
			name: "error-nil-ok",
			fn: func() {
				Ok(42, nil)
			},
		},
		{
			name: "true",
			fn: func() {
				True(true, "true")
			},
		},
		{
			name: "true-f",
			fn: func() {
				Truef(true, "true-%s", "f")
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
