package iter

import "github.com/fdschonborn/go-sandbox/option"

type filter[Item any] struct {
	iter Iterator[Item]
	f    func(Item) bool
}

func Filter[Item any](iter Iterator[Item], f func(Item) bool) Iterator[Item] {
	return &filter[Item]{
		iter: iter,
		f:    f,
	}
}

func (f *filter[Item]) Next() option.Option[Item] {
	for {
		item, ok := next(f.iter)
		if !ok {
			return option.None[Item]()
		}

		if f.f(item) {
			return option.Some(item)
		}
	}
}
