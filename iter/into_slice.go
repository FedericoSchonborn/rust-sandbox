package iter

import "github.com/fdschonborn/go-sandbox/option"

func IntoSlice[T any, Iter Iterator[T]](iter Iter) []T {
	var result []T

	// TODO: Replace this crap with something more appropriate.
	for {
		if option.Map(iter.Next(), func(value T) T {
			result = append(result, value)
			return value
		}).IsNone() {
			break
		}
	}

	return result
}
