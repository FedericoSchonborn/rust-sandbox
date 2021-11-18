package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/slices"
)

type mapIter[K comparable, V any] struct {
	m    map[K]V
	i    int
	keys []K
}

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func Iter[M ~map[K]V, K comparable, V any](m M) iter.Iterator[KeyValue[K, V]] {
	kvs := make([]KeyValue[K, V], len(m))

	var i int
	for key, value := range m {
		kvs[i] = KeyValue[K, V]{
			Key:   key,
			Value: value,
		}
		i++
	}

	return slices.Iter(kvs)
}
