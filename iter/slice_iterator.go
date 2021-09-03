package iter

import "github.com/fdschonborn/go-sandbox/option"

type sliceIterator[T any] struct {
	inner []T
	index int
}

func SliceIterator[T any](inner []T) Iterator[T] {
	return &sliceIterator[T]{
		inner: inner,
		index: 0,
	}
}

func (si *sliceIterator[T]) Next() option.Option[T] {
	if si.index >= len(si.inner) {
		return option.None[T]()
	}

	value := si.inner[si.index]
	si.index++
	return option.Some[T](value)
}
