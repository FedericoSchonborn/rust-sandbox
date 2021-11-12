package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
	"github.com/fdschonborn/go-sandbox/tuple"
)

type enumerate[Item any] struct {
	iter Iterator[Item]
	n    int
}

func Enumerate[Item any](iter Iterator[Item]) Iterator[tuple.Tuple2[int, Item]] {
	return &enumerate[Item]{
		iter: iter,
		n:    0,
	}
}

func (e *enumerate[Item]) Next() option.Option[tuple.Tuple2[int, Item]] {
	item, ok := next(e.iter)
	if !ok {
		return option.None[tuple.Tuple2[int, Item]]()
	}

	result := tuple.New2(e.n, item)
	e.n++
	return option.Some(result)
}
