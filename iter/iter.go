package iter

import (
	"github.com/fdschonborn/go-sandbox/cmp"
	"github.com/fdschonborn/go-sandbox/option"
	"github.com/fdschonborn/go-sandbox/zero"
)

type Iterator[Item any] interface {
	Next() option.Option[Item]
}

func next[Item any](iter Iterator[Item]) (item Item, ok bool) {
	if iter == nil {
		return zero.Zero[Item](), false
	}

	next := iter.Next()
	return next.UnwrapOrZero(), next.IsSome()
}

func Find[Item any](iter Iterator[Item], f func(Item) bool) option.Option[Item] {
	for {
		item, ok := next(iter)
		if !ok {
			return option.None[Item]()
		}

		if f(item) {
			return option.Some(item)
		}
	}
}

func FindMap[Item, Result any](iter Iterator[Item], f func(Item) option.Option[Result]) option.Option[Result] {
	return FilterMap(iter, f).Next()
}

func All[Item any](iter Iterator[Item], f func(Item) bool) bool {
	for {
		item, ok := next(iter)
		if !ok {
			break
		}

		if !f(item) {
			return false
		}
	}

	return true
}

func Any[Item any](iter Iterator[Item], f func(Item) bool) bool {
	for {
		item, ok := next(iter)
		if !ok {
			break
		}

		if f(item) {
			return true
		}
	}

	return false
}

func Collect[Item any](iter Iterator[Item]) []Item {
	var result []Item

	for {
		item, ok := next(iter)
		if !ok {
			break
		}

		result = append(result, item)
	}

	return result
}

func ForEach[Item any](iter Iterator[Item], f func(Item)) {
	for {
		item, ok := next(iter)
		if !ok {
			break
		}

		f(item)
	}
}

func Last[Item any](iter Iterator[Item]) option.Option[Item] {
	last := option.None[Item]()

	for {
		item, ok := next(iter)
		if !ok {
			return last
		}

		last = option.Some(item)
	}
}

func Fold[Item, Acc any](iter Iterator[Item], init Acc, f func(Acc, Item) Acc) Acc {
	acc := init

	for {
		item, ok := next(iter)
		if !ok {
			return acc
		}

		acc = f(acc, item)
	}
}

func Max[Item cmp.Ordered](iter Iterator[Item]) option.Option[Item] {
	max := option.None[Item]()
	for {
		item, ok := next(iter)
		if !ok {
			return max
		}

		if max.IsNone() || item > max.UnwrapUnchecked() {
			max = option.Some(item)
		}
	}
}

func Min[Item cmp.Ordered](iter Iterator[Item]) option.Option[Item] {
	min := option.None[Item]()
	for {
		item, ok := next(iter)
		if !ok {
			return min
		}

		if min.IsNone() || item < min.UnwrapUnchecked() {
			min = option.Some(item)
		}
	}
}
