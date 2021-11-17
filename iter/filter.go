package iter

import "github.com/fdschonborn/go-sandbox/zero"

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

func (f *filter[Item]) Next() (_ Item, ok bool) {
	for {
		item, ok := f.iter.Next()
		if !ok {
			return zero.Zero[Item](), false
		}

		if f.f(item) {
			return item, true
		}
	}
}
