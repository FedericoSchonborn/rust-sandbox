package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
	"github.com/fdschonborn/go-sandbox/tuple"
)

type enumerate[Item any, Iter Iterator[Item]] struct {
	iter Iter
	n    int
}

func Enumerate[Item any, Iter Iterator[Item]](iter Iter) Iterator[tuple.Tuple2[int, Item]] {
	return &enumerate[Item, Iter]{
		iter: iter,
		n:    0,
	}
}

func (e *enumerate[Item, Iter]) Next() option.Option[tuple.Tuple2[int, Item]] {
	item, ok := Next[Item](e.iter)
	if !ok {
		return option.None[tuple.Tuple2[int, Item]]()
	}

	result := tuple.New2(e.n, item)
	e.n++
	return option.Some(result)
}
