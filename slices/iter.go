package slices

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/zero"
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

func (si *sliceIter[T]) Next() (_ T, ok bool) {
	if si.i >= len(si.s) {
		return zero.Zero[T](), false
	}

	value := si.s[si.i]
	si.i++
	return value, true
}
