package slices

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type sliceIter[T any] struct {
	s []T
	i int
}

func Iter[T any](slice []T) iter.Iterator[T] {
	return &sliceIter[T]{
		s: slice,
		i: 0,
	}
}

func (si *sliceIter[T]) Next() option.Option[T] {
	if si.i >= len(si.s) {
		return option.None[T]()
	}
	defer func() { si.i++ }()

	return option.Some(si.s[si.i])
}
