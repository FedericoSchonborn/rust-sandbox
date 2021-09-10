package zero

type Defaulter[Self any] interface {
	Default() Self
}

func Default[T any]() T {
	var zero T
	if d, ok := (interface{})(zero).(Defaulter[T]); ok {
		return d.Default()
	}

	return zero
}

// Zero returns the zero value of type T.
func Zero[T any]() T {
	var zero T
	return zero
}

// IsZero reports wether a value of type T is the zero value of type T.
func IsZero[T comparable](value T) bool {
	return value == Zero[T]()
}
