package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type FilterFunc[T any] func(T) bool

type filter[T any] struct {
	iter Iterator[T]
	f    FilterFunc[T]
}

func Filter[T any](iter Iterator[T], f FilterFunc[T]) Iterator[T] {
	return &filter[T]{
		iter: iter,
		f:    f,
	}
}

func (f *filter[T]) Next() option.Option[T] {
	for {
		item, ok := f.iter.Next().Get()
		if !ok {
			return option.None[T]()
		}

		if f.f(item) {
			return option.Some(item)
		}
	}
}
