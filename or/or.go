package or

import "github.com/fdschonborn/go-sandbox/zero"

func Or[T comparable](value T, defaults ...T) T {
	zero := zero.Zero[T]()
	if value != zero {
		return value
	}

	for _, value := range defaults {
		if value != zero {
			return value
		}
	}

	return zero
}
