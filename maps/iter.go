package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
)

type mapIter[K comparable, V any] struct {
	m    map[K]V
	i    int
	keys []K
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func Iter[K comparable, V any](inner map[K]V) iter.Iterator[Pair[K, V]] {
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

func (mi *mapIter[K, V]) Next() (_ Pair[K, V], ok bool) {
	if mi.i >= len(mi.keys) {
		return Pair[K, V]{}, false
	}
	defer func() { mi.i++ }()

	return Pair[K, V]{
		Key:   mi.keys[mi.i],
		Value: mi.m[mi.keys[mi.i]],
	}, true
}
