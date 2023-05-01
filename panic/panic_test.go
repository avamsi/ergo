package panic

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
			name: "assert-false",
			fn: func() {
				Assert(false, "assert")
			},
			want: "assert",
		},
		{
			name: "assertf-false",
			fn: func() {
				Assertf(false, "assert%s", "f")
			},
			want: "assertf",
		},
		{
			name: "must0-err",
			fn: func() {
				Must0(errors.New("must0"))
			},
			want: "must0",
		},
		{
			name: "must1-err",
			fn: func() {
				Must1(1, errors.New("must1"))
			},
			want: "must1",
		},
		{
			name: "must2-err",
			fn: func() {
				Must2(1, 2, errors.New("must2"))
			},
			want: "must2",
		},
		{
			name: "must3-err",
			fn: func() {
				Must3(1, 2, 3, errors.New("must3"))
			},
			want: "must3",
		},
		{
			name: "panicf",
			fn: func() {
				Panicf("panic%s", "f")
			},
			want: "panicf",
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
			name: "assert-true",
			fn: func() {
				Assert(true, "assert-true")
			},
		},
		{
			name: "assertf-true",
			fn: func() {
				Assertf(true, "assertf-%t", true)
			},
		},
		{
			name: "must0-nil",
			fn: func() {
				Must0(nil)
			},
		},
		{
			name: "must1-nil",
			fn: func() {
				Must1(errors.New("1"), nil)
			},
		},
		{
			name: "must2-nil",
			fn: func() {
				Must2(1, errors.New("2"), nil)
			},
		},
		{
			name: "must3-nil",
			fn: func() {
				Must3(1, 2, errors.New("3"), nil)
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
