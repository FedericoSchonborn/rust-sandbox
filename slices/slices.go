package slices

func Swap[T any](s []T, a, b int) {
	s[a], s[b] = s[b], s[a]
}

func Repeat[T any](s []T, n int) []T {
	new := make([]T, len(s)*n)

	var pos int
	for rem := 0; rem < n; rem++ {
		for idx := 0; idx < len(s); idx++ {
			new[pos] = s[idx]
			pos++
		}
	}

	return new
}
