//go:build broken

package slices

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type sliceIter[S ~[]T, T any] struct {
	s S
	i int
}

func Iter[S ~[]T, T any](slice S) iter.Iterator[T] {
	return &sliceIter[S, T]{
		s: slice,
		i: 0,
	}
}

func (si *sliceIter[S, T]) Next() option.Option[T] {
	if si.i >= len(si.s) {
		return option.None[T]()
	}
	defer func() { si.i++ }()

	return option.Some(si.s[si.i])
}
