package pair_test

import (
	"testing"

	"github.com/avamsi/ergo/pair"
	"github.com/google/go-cmp/cmp"
)

func TestPair(t *testing.T) {
	{
		var (
			first      = 1
			second     = 2
			got1, got2 = pair.New(first, second).Unpack()
		)
		if !cmp.Equal(got1, first) || !cmp.Equal(got2, second) {
			t.Errorf("New(%d, %d).Unpack() = %d, %d, want %[1]d, %d", first, second, got1, got2)
		}
	}
	{
		var (
			first      map[int]int
			second     error
			got1, got2 = pair.New(first, second).Unpack()
		)
		if !cmp.Equal(got1, first) || !cmp.Equal(got2, second) {
			t.Errorf("New(%d, %d).Unpack() = %d, %d, want %[1]d, %d", first, second, got1, got2)
		}
	}
}
