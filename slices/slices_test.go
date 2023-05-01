package slices

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunks(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want [][]int
	}{
		{
			name: "nil",
			s:    nil,
			n:    0,
			want: nil,
		},
		{
			name: "nils",
			s:    nil,
			n:    3,
			want: [][]int{nil, nil, nil},
		},
		{
			name: "empty",
			s:    []int{},
			n:    5,
			want: [][]int{{}, {}, {}, {}, {}},
		},
		{
			name: "non-empty",
			s:    []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
		},
		{
			name: "non-empty-over-subscribed",
			s:    []int{1, 2, 3},
			n:    5,
			want: [][]int{{1}, {2}, {3}, {}, {}},
		},
		{
			name: "non-empty-optimal",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:    4,
			// Should avoid both 2-2-2-4 and 3-3-3-1 chunks.
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8},
				{9, 10},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got [][]int
			Chunks(test.s, test.n, func(i int, chunk []int) {
				if l := len(got); i != l {
					t.Errorf("(i=)%v != %v(=len(%#v(=got))\n", i, l, got)
				}
				got = append(got, chunk)
			})
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Chunks(%#v) returned diff(-want +got):\n%v", test.s, diff)
			}
		})
	}
}
