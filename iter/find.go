package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type FindFunc[T any] func(T) bool

func Find[T any](iter Iterator[T], f FindFunc[T]) option.Option[T] {
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			return option.None[T]()
		}

		if f(item) {
			return option.Some(item)
		}
	}
}
