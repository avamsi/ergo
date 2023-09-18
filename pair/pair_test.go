package pair

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPair(t *testing.T) {
	{
		var (
			first      = 1
			second     = 2
			got1, got2 = New(first, second).Unpack()
		)
		if !cmp.Equal(got1, first) || !cmp.Equal(got2, second) {
			t.Errorf("New(%v, %v).Unpack() = %v, %v, want %[1]v, %v", first, second, got1, got2)
		}
	}
	{
		var (
			first      map[int]int
			second     error
			got1, got2 = New(first, second).Unpack()
		)
		if !cmp.Equal(got1, first) || !cmp.Equal(got2, second) {
			t.Errorf("New(%v, %v).Unpack() = %v, %v, want %[1]v, %v", first, second, got1, got2)
		}
	}
}
