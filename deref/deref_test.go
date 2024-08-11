package deref_test

import (
	"testing"

	"github.com/avamsi/ergo/deref"
)

func TestOr(t *testing.T) {
	tests := []struct {
		name  string
		ptr   *int
		value int
		want  int
	}{
		{
			name:  "nil",
			ptr:   nil,
			value: 42,
			want:  42,
		},
		{
			name:  "non-nil",
			ptr:   new(int),
			value: 42,
			want:  0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := deref.Or(test.ptr, test.value); got != test.want {
				t.Errorf("Or(%#v, %#v) = %#v, want %#v", test.ptr, test.value, got, test.want)
			}
		})
	}
}
