package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type FilterMapFunc[T, U any] func(T) option.Option[U]

type filterMap[T, U any] struct {
	iter Iterator[T]
	f    FilterMapFunc[T, U]
}

func FilterMap[T, U any](iter Iterator[T], f FilterMapFunc[T, U]) Iterator[U] {
	return &filterMap[T, U]{
		iter: iter,
		f:    f,
	}
}

func (fm *filterMap[T, U]) Next() option.Option[U] {
	for {
		item, ok := fm.iter.Next().Get()
		if !ok {
			return option.None[U]()
		}

		result, ok := fm.f(item).Get()
		if !ok {
			continue
		}

		return option.Some(result)
	}
}
