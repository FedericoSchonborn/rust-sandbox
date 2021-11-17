package iter

import (
	"github.com/fdschonborn/go-sandbox/zero"
)

type filterMap[Item, Result any] struct {
	iter Iterator[Item]
	f    func(Item) (Result, bool)
}

func FilterMap[Item, Result any](iter Iterator[Item], f func(Item) (Result, bool)) Iterator[Result] {
	return &filterMap[Item, Result]{
		iter: iter,
		f:    f,
	}
}

func (fm *filterMap[Item, Result]) Next() (_ Result, ok bool) {
	for {
		item, ok := fm.iter.Next()
		if !ok {
			return zero.Zero[Result](), false
		}

		result, ok := fm.f(item)
		if !ok {
			continue
		}

		return result, true
	}
}
