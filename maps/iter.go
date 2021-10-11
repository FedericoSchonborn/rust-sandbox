package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
	"github.com/fdschonborn/go-sandbox/tuple"
)

type mapIter[K comparable, V any] struct {
	m    map[K]V
	i    int
	keys []K
}

func Iter[K comparable, V any](inner map[K]V) iter.Iterator[tuple.Tuple2[K, V]] {
	keys := make([]K, len(inner))

	var i int
	for key := range inner {
		keys[i] = key
		i++
	}

	return &mapIter[K, V]{
		m:    inner,
		i:    0,
		keys: keys,
	}
}

func (mi *mapIter[K, V]) Next() option.Option[tuple.Tuple2[K, V]] {
	if mi.i >= len(mi.keys) {
		return option.None[tuple.Tuple2[K, V]]()
	}

	key := mi.keys[mi.i]
	value := mi.m[key]
	mi.i++
	return option.Some(tuple.New2[K, V](key, value))
}
