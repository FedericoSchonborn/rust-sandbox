package iter

import "github.com/fdschonborn/go-sandbox/option"

type map_[Item, Result any, Iter Iterator[Item]] struct {
	iter Iter
	f    func(Item) Result
}

func Map[Item, Result any, Iter Iterator[Item]](iter Iter, f func(Item) Result) Iterator[Result] {
	return &map_[Item, Result, Iter]{
		iter: iter,
		f:    f,
	}
}

func (m *map_[Item, Result, Iter]) Next() option.Option[Result] {
	return option.Map(m.iter.Next(), m.f)
}
