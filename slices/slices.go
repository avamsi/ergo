package slices

// Chunks breaks the given slice into n slices of (almost) the same size.
func Chunks[E any](s []E, n int, yield func(int, []E)) {
	if n == 0 {
		return
	}
	var (
		size, reminder = len(s) / n, len(s) % n
		start, end     int
	)
	for i := 0; i < n; i++ {
		start, end = end, end+size
		if reminder > 0 {
			reminder--
			end++
		}
		yield(i, s[start:end])
	}
}
