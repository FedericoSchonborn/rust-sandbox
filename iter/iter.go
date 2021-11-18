package iter

import (
	"constraints"

	"github.com/fdschonborn/go-sandbox/option"
)

type Iterator[T any] interface {
	Next() option.Option[T]
}

func Find[T any](iter Iterator[T], f func(T) bool) option.Option[T] {
	for {
		item, ok := iter.Next().Get()
		if !ok {
			return option.None[T]()
		}

		if f(item) {
			return option.Some(item)
		}
	}
}

type FindMapFunc[T, U any] func(T) option.Option[U]

func FindMap[T, U any](iter Iterator[T], f FindMapFunc[T, U]) option.Option[U] {
	// The func is a fugly workaround because FilterMapFunc != FindMapFunc
	return FilterMap(iter, func(item T) option.Option[U] { return f(item) }).Next()
}

func All[T any](iter Iterator[T], f func(T) bool) bool {
	for {
		item, ok := iter.Next().Get()
		if !ok {
			break
		}

		if !f(item) {
			return false
		}
	}

	return true
}

func Any[T any](iter Iterator[T], f func(T) bool) bool {
	for {
		item, ok := iter.Next().Get()
		if !ok {
			break
		}

		if f(item) {
			return true
		}
	}

	return false
}

func Collect[T any](iter Iterator[T]) []T {
	var result []T

	for {
		item, ok := iter.Next().Get()
		if !ok {
			break
		}

		result = append(result, item)
	}

	return result
}

func ForEach[T any](iter Iterator[T], f func(T)) {
	for {
		item, ok := iter.Next().Get()
		if !ok {
			break
		}

		f(item)
	}
}

func Last[T any](iter Iterator[T]) option.Option[T] {
	last := option.None[T]()

	for {
		item, ok := iter.Next().Get()
		if !ok {
			return last
		}

		last = option.Some(item)
	}
}

type FoldFunc[T, A any] func(A, T) A

func Fold[T, A any](iter Iterator[T], init A, f FoldFunc[T, A]) A {
	acc := init

	for {
		item, ok := iter.Next().Get()
		if !ok {
			return acc
		}

		acc = f(acc, item)
	}
}

func Max[T constraints.Ordered](iter Iterator[T]) option.Option[T] {
	max := option.None[T]()
	for {
		item, ok := iter.Next().Get()
		if !ok {
			return max
		}

		if !ok || item > max.UnwrapOrZero() {
			max = option.Some(item)
		}
	}
}

func Min[T constraints.Ordered](iter Iterator[T]) option.Option[T] {
	min := option.None[T]()
	for {
		item, ok := iter.Next().Get()
		if !ok {
			return min
		}

		if !ok || item < min.UnwrapOrZero() {
			min = option.Some(item)
		}
	}
}
