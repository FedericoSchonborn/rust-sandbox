package iter

import "github.com/fdschonborn/go-sandbox/option"

type filterMap[Item, Result any, Iter Iterator[Item]] struct {
	iter Iter
	f    func(Item) option.Option[Result]
}

func FilterMap[Item, Result any, Iter Iterator[Item]](iter Iter, f func(Item) option.Option[Result]) Iterator[Result] {
	return &filterMap[Item, Result, Iter]{
		iter: iter,
		f:    f,
	}
}

func (fm *filterMap[Item, Result, Iter]) Next() option.Option[Result] {
	for {
		item, ok := Next[Item, Iter](fm.iter)
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
