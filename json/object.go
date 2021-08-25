package json

import (
	"errors"
	"reflect"
)

type Object[T any] map[string]T

// TODO: Support composite types.
func (o Object[T]) Get(field string, out *T) error {
	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr {
		return errors.New("pointer required")
	}
	outElem := outValue.Elem()

	value, ok := o[field]
	if !ok {
		return errors.New("not found")
	}

	valueValue := reflect.ValueOf(value)
	if valueValue.Type() != outElem.Type() {
		return errors.New("type mismatch")
	}

	outElem.Set(valueValue)
	return nil
}
