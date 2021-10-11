package json

import (
	"errors"

	"github.com/fdschonborn/go-sandbox/constraints"
	"github.com/fdschonborn/go-sandbox/zero"
)

type Object[K constraints.String, V any] map[K]V

// TODO: Support composite types.
func (o Object[K, V]) Get(field K) (V, error) {
	value, ok := o[field]
	if !ok {
		return zero.Zero[V](), errors.New("not found")
	}

	return value, nil
}
