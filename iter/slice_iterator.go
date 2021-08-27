package iter

import "github.com/fdschonborn/go-sandbox/option"

type SliceIterator[T any] struct {
	inner []T
	index int
}

func NewSliceIterator[T any](inner []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		inner: inner,
		index: 0,
	}
}

func (si *SliceIterator[T]) Next() option.Option[T] {
	if si.index >= len(si.inner) {
		return option.None[T]()
	}
	si.index++

	return option.Some[T](si.inner[si.index])
}
