//go:build broken

package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
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

	return iter.Slice(values)
}
