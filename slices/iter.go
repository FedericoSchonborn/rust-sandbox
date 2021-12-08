package slices

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type sliceIter[S ~[]T, T any] struct {
	s S
	l int
	i int
}

func Iter[S ~[]T, T any](s S) iter.Iterator[T] {
	return &sliceIter[S, T]{
		s: s,
		l: len(s),
		i: 0,
	}
}

func (si *sliceIter[S, T]) Next() option.Option[T] {
	if si.i >= si.l {
		return option.None[T]()
	}
	defer func() { si.i++ }()

	return option.Some(si.s[si.i])
}
