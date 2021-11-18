package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type Enumerated[Item any] struct {
	N    int
	Item Item
}

type enumerate[T any] struct {
	iter Iterator[T]
	n    int
}

func Enumerate[T any](iter Iterator[T]) Iterator[Enumerated[T]] {
	return &enumerate[T]{
		iter: iter,
		n:    0,
	}
}

func (e *enumerate[T]) Next() option.Option[Enumerated[T]] {
	orig, ok := e.iter.Next().Get()
	if !ok {
		return option.None[Enumerated[T]]()
	}
	defer func() { e.n++ }()

	return option.Some(Enumerated[T]{N: e.n, Item: orig})
}
