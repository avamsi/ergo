package group_test

import (
	stdcmp "cmp"
	"testing"

	"github.com/avamsi/ergo/group"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCollector(t *testing.T) {
	tests := []struct {
		name string
		fs   []func(func(int))
		want []int
	}{
		{
			name: "nil",
			fs:   nil,
			want: nil,
		},
		{
			name: "1-3",
			fs: []func(func(int)){
				func(collect func(int)) { collect(1) },
				func(collect func(int)) { collect(2) },
				func(collect func(int)) { collect(3) },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := group.NewCollector(make(chan int))
			for _, f := range test.fs {
				c.Go(func() { f(c.Collect) })
			}
			var got []int
			for i := range c.Close() {
				got = append(got, i)
			}
			sortOpt := cmpopts.SortSlices(stdcmp.Less[int])
			if diff := cmp.Diff(test.want, got, sortOpt); diff != "" {
				t.Errorf("diff(-want +got): %s", diff)
			}
		})
	}
}
