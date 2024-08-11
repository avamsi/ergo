package slices

import (
	"iter"

	"github.com/avamsi/ergo/assert"
)

// Shard breaks the given slice into n slices of (almost) the same size.
func Shard[S ~[]E, E any](s S, n int) iter.Seq[S] {
	assert.Truef(n > 0, "n (%d) <= 0", n)
	return func(yield func(S) bool) {
		var (
			size, remainder = len(s) / n, len(s) % n
			start, end      int
		)
		for range n {
			start, end = end, end+size
			if remainder > 0 {
				remainder--
				end++
			}
			if !yield(s[start:end:end]) {
				return
			}
		}
	}
}
