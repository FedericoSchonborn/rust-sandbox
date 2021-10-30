package zero

// Zero returns the zero value of type T.
func Zero[T any]() T {
	// TODO: If T is interface{}, resolve to the original type.
	var zero T
	return zero
}

// IsZero reports wether a value of type T is the zero value of type T.
func IsZero[T comparable](value T) bool {
	return value == Zero[T]()
}
