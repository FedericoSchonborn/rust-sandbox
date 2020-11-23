package typeof

import "reflect"

func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).Name()
}
