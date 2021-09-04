package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
	"github.com/fdschonborn/go-sandbox/tuple"
)

type fromMap[K comparable, V any] struct {
	inner map[K]V
	keys  []K
	index int
}

func FromMap[K comparable, V any](inner map[K]V) Iterator[tuple.Tuple2[K, V]] {
	keys := make([]K, len(inner))
	index := 0
	for key := range inner {
		keys[index] = key
		index++
	}

	return &fromMap[K, V]{
		inner: inner,
		keys:  keys,
		index: 0,
	}
}

func (mi *fromMap[K, V]) Next() option.Option[tuple.Tuple2[K, V]] {
	if mi.index >= len(mi.keys) {
		return option.None[tuple.Tuple2[K, V]]()
	}

	key := mi.keys[mi.index]
	value := mi.inner[key]
	mi.index++
	return option.Some(tuple.NewTuple2[K, V](key, value))
}
