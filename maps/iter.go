package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type Paired[K comparable, V any] struct {
	K K
	V V
}

type mapIter[M ~map[K]V, K comparable, V any] struct {
	m M
	k []K
	l int
	i int
}

func Iter[M ~map[K]V, K comparable, V any](m M) iter.Iterator[Paired[K, V]] {
	k := make([]K, len(m))
	var n int
	for key := range m {
		k[n] = key
		n++
	}

	return &mapIter[M, K, V]{
		m: m,
		k: k,
		l: len(k),
		i: 0,
	}
}

func (mi *mapIter[M, K, V]) Next() option.Option[Paired[K, V]] {
	if mi.i >= mi.l {
		return option.None[Paired[K, V]]()
	}
	defer func() { mi.i++ }()

	key := mi.k[mi.i]
	value := mi.m[key]
	return option.Some(Paired[K, V]{K: key, V: value})
}
