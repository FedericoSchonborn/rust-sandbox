package or

import "reflect"

func Or(v, d interface{}) interface{} {
	value := reflect.ValueOf(v)
	dType := reflect.TypeOf(d)

	if dType != value.Type() {
		panic("Type mismatch")
	}

	if !value.IsZero() {
		return v
	}

	return d
}
