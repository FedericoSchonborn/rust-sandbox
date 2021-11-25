package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type sliceIterator[T any, S ~[]T] struct {
	s S
	i int
}

func Slice[T any, S ~[]T](slice S) Iterator[T] {
	return &sliceIterator[T, S]{
		s: slice,
		i: 0,
	}
}

func (si *sliceIterator[T, S]) Next() option.Option[T] {
	if si.i >= len(si.s) {
		return option.None[T]()
	}
	defer func() { si.i++ }()

	return option.Some(si.s[si.i])
}
