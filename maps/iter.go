package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

type mapIter[K comparable, V any] struct {
	m    map[K]V
	i    int
	keys []K
}

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

func Iter[K comparable, V any](inner map[K]V) iter.Iterator[Item[K, V]] {
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

func (mi *mapIter[K, V]) Next() option.Option[Item[K, V]] {
	if mi.i >= len(mi.keys) {
		return option.None[Item[K, V]]()
	}
	defer func() { mi.i++ }()

	return option.Some(Item[K, V]{Key: mi.keys[mi.i], Value: mi.m[mi.keys[mi.i]]})
}
