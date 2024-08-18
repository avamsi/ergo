package slices_test

import (
	"testing"

	"github.com/avamsi/ergo/slices"
	"github.com/google/go-cmp/cmp"
)

func TestShard(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want [][]int
	}{
		{
			name: "nil",
			s:    nil,
			n:    1,
			want: [][]int{nil},
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
			// Should avoid both 2-2-2-4 and 3-3-3-1 shards.
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
			for shard := range slices.Shard(test.s, test.n) {
				got = append(got, shard)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Shard(%d) has diff(-want +got):\n%s", test.s, diff)
			}
		})
	}
}
