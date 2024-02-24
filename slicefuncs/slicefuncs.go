package slicefuncs

// Map transforms the elements of a slice and returns new slice of the same length.
func Map[T any, V any](in []T, transform func(T) V) []V {
	var out []V

	for _, v := range in {
		out = append(out, transform(v))
	}
	return out
}

// Filter either discards the elements or keeps them if the keep function returns true.
// It returns a new slice.
func Filter[T any](in []T, keep func(T) bool) []T {
	var out []T

	for _, v := range in {
		if keep(v) {
			out = append(out, v)
		}

	}
	return out
}
