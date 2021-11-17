package iter

import (
	"github.com/fdschonborn/go-sandbox/cmp"
	"github.com/fdschonborn/go-sandbox/zero"
)

type Iterator[Item any] interface {
	Next() (item Item, ok bool)
}

func Find[Item any](iter Iterator[Item], f func(Item) bool) (_ Item, ok bool) {
	for {
		item, ok := iter.Next()
		if !ok {
			return zero.Zero[Item](), false
		}

		if f(item) {
			return item, true
		}
	}
}

func FindMap[Item, Result any](iter Iterator[Item], f func(Item) (Result, bool)) (_ Result, ok bool) {
	return FilterMap(iter, f).Next()
}

func All[Item any](iter Iterator[Item], f func(Item) bool) bool {
	for {
		item, ok := iter.Next()
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
		item, ok := iter.Next()
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
		item, ok := iter.Next()
		if !ok {
			break
		}

		result = append(result, item)
	}

	return result
}

func ForEach[Item any](iter Iterator[Item], f func(Item)) {
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		f(item)
	}
}

func Last[Item any](iter Iterator[Item]) (_ Item, ok bool) {
	last, _ok := zero.Zero[Item](), false

	for {
		item, ok := iter.Next()
		if !ok {
			return last, _ok
		}

		last, _ok = item, true
	}
}

func Fold[Item, Acc any](iter Iterator[Item], init Acc, f func(Acc, Item) Acc) Acc {
	acc := init

	for {
		item, ok := iter.Next()
		if !ok {
			return acc
		}

		acc = f(acc, item)
	}
}

func Max[Item cmp.Ordered](iter Iterator[Item]) (_ Item, ok bool) {
	max := zero.Zero[Item]()
	for {
		item, ok := iter.Next()
		if !ok {
			return max, true
		}

		if !ok || item > max {
			max = item
		}
	}
}

func Min[Item cmp.Ordered](iter Iterator[Item]) (_ Item, ok bool) {
	min := zero.Zero[Item]()
	for {
		item, ok := iter.Next()
		if !ok {
			return min, true
		}

		if !ok || item < min {
			min = item
		}
	}
}
