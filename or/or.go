package or

import "reflect"

func Or(v interface{}, defaults ...interface{}) interface{} {
	value := reflect.ValueOf(v)

	if !value.IsZero() {
		return v
	}

	for _, d := range defaults {
		dvalue := reflect.ValueOf(d)
		if dvalue.Type() != value.Type() {
			panic("Type mismatch")
		}

		if !dvalue.IsZero() {
			return d
		}
	}

	return nil
}
