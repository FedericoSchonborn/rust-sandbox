package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/slices"
)

type valuesIter[V any] struct {
	values []V
}

func Values[M ~map[K]V, K comparable, V any](m M) iter.Iterator[V] {
	values := make([]V, len(m))

	var i int
	for _, value := range m {
		values[i] = value
		i++
	}

	return slices.Iter(values)
}
