package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type Iterator[Item any] interface {
	Next() option.Option[Item]
}

func Next[Item any, Iter Iterator[Item]](iter Iter) (item Item, ok bool) {
	next := iter.Next()
	return next.UnwrapOrZero(), next.IsSome()
}

func All[Item any, Iter Iterator[Item]](iter Iter, f func(Item) bool) bool {
	for {
		item, ok := Next[Item, Iter](iter)
		if !ok {
			break
		}

		if !f(item) {
			return false
		}
	}

	return true
}

func Any[Item any, Iter Iterator[Item]](iter Iter, f func(Item) bool) bool {
	for {
		item, ok := Next[Item, Iter](iter)
		if !ok {
			break
		}

		if f(item) {
			return true
		}
	}

	return false
}

func Collect[Item any, Iter Iterator[Item]](iter Iter) []Item {
	var result []Item

	for {
		item, ok := Next[Item, Iter](iter)
		if !ok {
			break
		}

		result = append(result, item)
	}

	return result
}

func ForEach[Item any, Iter Iterator[Item]](iter Iter, f func(Item)) {
	for {
		item, ok := Next[Item, Iter](iter)
		if !ok {
			break
		}

		f(item)
	}
}

func Last[Item any, Iter Iterator[Item]](iter Iter) option.Option[Item] {
	last := option.None[Item]()
	for {
		item, ok := Next[Item, Iter](iter)
		if !ok {
			return last
		}

		last = option.Some(item)
	}
}
