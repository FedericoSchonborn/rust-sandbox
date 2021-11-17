package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type enumerate[Item any] struct {
	iter Iterator[Item]
	n    int
}

type EnumerateItem[Item any] struct {
	N    int
	Item Item
}

func Enumerate[Item any](iter Iterator[Item]) Iterator[EnumerateItem[Item]] {
	return &enumerate[Item]{
		iter: iter,
		n:    0,
	}
}

func (e *enumerate[Item]) Next() option.Option[EnumerateItem[Item]] {
	item, ok := next(e.iter)
	if !ok {
		return option.None[EnumerateItem[Item]]()
	}
	defer func() { e.n++ }()

	return option.Some(EnumerateItem[Item]{N: e.n, Item: item})
}
