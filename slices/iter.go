package slices

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type sliceIter[T any] struct {
	s []T
	i int
}

func Iter[T any](s []T) iter.Iterator[T] {
	return &sliceIter[T]{
		s: s,
		i: 0,
	}
}

func (si *sliceIter[T]) Next() option.Option[T] {
	if si.i >= len(si.s) {
		return option.None[T]()
	}

	value := si.s[si.i]
	si.i++
	return option.Some(value)
}
