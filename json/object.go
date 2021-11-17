package json

import (
	"errors"

	"github.com/fdschonborn/go-sandbox/zero"
)

type Object[V any] map[string]V

// TODO: Support composite types.
func (o Object[T]) Get(key string) (T, error) {
	value, ok := o[key]
	if !ok {
		return zero.Zero[T](), errors.New("not found")
	}

	return value, nil
}
