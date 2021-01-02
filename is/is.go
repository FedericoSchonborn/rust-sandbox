package is

import (
	"math"
	"reflect"
)

func Even(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch kind := v.Kind(); kind {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return v.Int()%2 == 0
	case reflect.Float32,
		reflect.Float64:
		return math.Mod(v.Float(), 2) == 0
	default:
		return false
	}
}

func Odd(i interface{}) bool {
	return !Even(i)
}
