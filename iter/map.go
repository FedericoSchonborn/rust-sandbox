package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type MapFunc[T, U any] func(T) U

type _map[T, U any] struct {
	iter Iterator[T]
	f    MapFunc[T, U]
}

func Map[T, U any](iter Iterator[T], f MapFunc[T, U]) Iterator[U] {
	return &_map[T, U]{
		iter: iter,
		f:    f,
	}
}

func (m *_map[T, U]) Next() option.Option[U] {
	item, ok := m.iter.Next().Get()
	if !ok {
		return option.None[U]()
	}

	return option.Some(m.f(item))
}
