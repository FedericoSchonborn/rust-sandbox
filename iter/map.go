package iter

import "github.com/fdschonborn/go-sandbox/option"

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

func (m *map_[Item, Result]) Next() option.Option[Result] {
	return option.Map(m.iter.Next(), m.f)
}
