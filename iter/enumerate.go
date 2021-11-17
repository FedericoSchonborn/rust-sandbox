package iter

import (
	"github.com/fdschonborn/go-sandbox/zero"
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

func (e *enumerate[Item]) Next() (_ EnumerateItem[Item], ok bool) {
	orig, ok := e.iter.Next()
	if !ok {
		return zero.Zero[EnumerateItem[Item]](), false
	}
	defer func() { e.n++ }()

	return EnumerateItem[Item]{N: e.n, Item: orig}, true
}
