package iter

import "github.com/fdschonborn/go-sandbox/option"

type filterMap[Item, Result any] struct {
	iter Iterator[Item]
	f    func(Item) option.Option[Result]
}

func FilterMap[Item, Result any](iter Iterator[Item], f func(Item) option.Option[Result]) Iterator[Result] {
	return &filterMap[Item, Result]{
		iter: iter,
		f:    f,
	}
}

func (fm *filterMap[Item, Result]) Next() option.Option[Result] {
	for {
		item, ok := next(fm.iter)
		if !ok {
			return option.None[Result]()
		}

		result := fm.f(item)
		if result.IsNone() {
			continue
		}

		return result
	}
}
