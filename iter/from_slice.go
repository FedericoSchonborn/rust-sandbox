package iter

import "github.com/fdschonborn/go-sandbox/option"

type fromSlice[T any] struct {
	inner []T
	index int
}

func FromSlice[T any](inner []T) Iterator[T] {
	return &fromSlice[T]{
		inner: inner,
		index: 0,
	}
}

func (si *fromSlice[T]) Next() option.Option[T] {
	if si.index >= len(si.inner) {
		return option.None[T]()
	}

	value := si.inner[si.index]
	si.index++
	return option.Some[T](value)
}
