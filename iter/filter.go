package iter

import "github.com/fdschonborn/go-sandbox/option"

type filter[Item any, Iter Iterator[Item]] struct {
	iter Iter
	f    func(Item) bool
}

func Filter[Item any, Iter Iterator[Item]](iter Iter, f func(Item) bool) Iterator[Item] {
	return &filter[Item, Iter]{
		iter: iter,
		f:    f,
	}
}

func (f *filter[Item, Iter]) Next() option.Option[Item] {
	for {
		item, ok := Next[Item](f.iter)
		if !ok {
			return option.None[Item]()
		}

		if f.f(item) {
			return option.Some(item)
		}
	}
}
