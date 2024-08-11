package iter_test

import (
	stditer "iter"
	"testing"

	"github.com/avamsi/ergo/iter"
	"github.com/google/go-cmp/cmp"
)

type kv struct {
	K, V int
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		name string
		s    stditer.Seq[int]
		want []kv
	}{
		{
			name: "empty",
			s:    func(func(int) bool) {},
			want: nil,
		},
		{
			name: "1-3",
			s: func(yield func(int) bool) {
				_ = yield(1) && yield(2) && yield(3)
			},
			want: []kv{{0, 1}, {1, 2}, {2, 3}},
		},
		{
			name: "5-3",
			s: func(yield func(int) bool) {
				_ = yield(5) && yield(4) && yield(3)
			},
			want: []kv{{0, 5}, {1, 4}, {2, 3}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got []kv
			for i, v := range iter.Enumerate(test.s) {
				got = append(got, kv{i, v})
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("diff(-want +got):\n%s", diff)
			}
		})
	}
}
