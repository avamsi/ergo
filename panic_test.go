package ergo_test

import (
	"testing"

	"github.com/avamsi/ergo"
)

func TestPanic(t *testing.T) {
	tests := []struct {
		name string
		f    func()
		want string
	}{
		{
			name: "nil",
			f: func() {
				ergo.Panic(nil)
			},
			want: "<nil>",
		},
		{
			name: "empty",
			f: func() {
				ergo.Panic("")
			},
			want: "",
		},
		{
			name: "1-3",
			f: func() {
				ergo.Panic(1, 2, 3)
			},
			want: "1 2 3",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if got := recover(); got != test.want {
					t.Errorf("got panic %#v, want %#v", test.want, got)
				}
			}()
			test.f()
		})
	}
}

func TestPanicf(t *testing.T) {
	tests := []struct {
		name string
		f    func()
		want string
	}{
		{
			name: "empty",
			f: func() {
				ergo.Panicf("")
			},
			want: "",
		},
		{
			name: "1+2",
			f: func() {
				ergo.Panicf("%d + %d != %d", 1, 2, 4)
			},
			want: "1 + 2 != 4",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if got := recover(); got != test.want {
					t.Errorf("got panic %#v, want %#v", test.want, got)
				}
			}()
			test.f()
		})
	}
}
