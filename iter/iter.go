package iter

import "iter"

func Enumerate[V any](s iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		var i int
		for v := range s {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}
