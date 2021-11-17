package iter

import (
	"github.com/fdschonborn/go-sandbox/zero"
)

type map_[Item, Result any] struct {
	iter Iterator[Item]
	f    func(Item) Result
}

func Map[Item, Result any](iter Iterator[Item], f func(Item) Result) Iterator[Result] {
	return &map_[Item, Result]{
		iter: iter,
		f:    f,
	}
}

func (m *map_[Item, Result]) Next() (_ Result, ok bool) {
	item, ok := m.iter.Next()
	if !ok {
		return zero.Zero[Result](), false
	}

	return m.f(item), true
}
