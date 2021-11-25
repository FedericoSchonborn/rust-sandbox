//go:build broken

package maps

import (
	"github.com/fdschonborn/go-sandbox/iter"
)

type keysIter[K comparable] struct {
	inner []K
}

func Keys[M ~map[K]V, K comparable, V any](m M) iter.Iterator[K] {
	keys := make([]K, len(m))

	var i int
	for key := range m {
		keys[i] = key
		i++
	}

	return iter.Slice(keys)
}
