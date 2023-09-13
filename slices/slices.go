package slices

import "github.com/avamsi/ergo/assert"

// Shard breaks the given slice into n slices of (almost) the same size.
func Shard[E any](s []E, n int) func(func(int, []E)) {
	assert.Truef(n > 0, "n (%d) <= 0", n)
	return func(yield func(int, []E)) {
		var (
			size, remainder = len(s) / n, len(s) % n
			start, end      int
		)
		for i := 0; i < n; i++ {
			start, end = end, end+size
			if remainder > 0 {
				remainder--
				end++
			}
			yield(i, s[start:end])
		}
	}
}
