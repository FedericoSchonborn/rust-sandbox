package zero

// Zero returns the zero value of type T.
func Zero[T any]() T {
	var zero T
	return zero
}
