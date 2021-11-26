package iter

import (
	"constraints"

	"github.com/fdschonborn/go-sandbox/option"
)

type Iterator[T any] interface {
	Next() option.Option[T]
}

func Collect[T any](iter Iterator[T]) []T {
	var result []T

	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			break
		}

		result = append(result, item)
	}

	return result
}

func Last[T any](iter Iterator[T]) option.Option[T] {
	last := option.None[T]()

	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			return last
		}

		last = option.Some(item)
	}
}

func Max[T constraints.Ordered](iter Iterator[T]) option.Option[T] {
	max := option.None[T]()
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			return max
		}

		if n, mok := max.UnwrapOrZero(); !ok || (mok && item > n) {
			max = option.Some(item)
		}
	}
}

func Min[T constraints.Ordered](iter Iterator[T]) option.Option[T] {
	min := option.None[T]()
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			return min
		}

		if n, mok := min.UnwrapOrZero(); !ok || (mok && item < n) {
			min = option.Some(item)
		}
	}
}
