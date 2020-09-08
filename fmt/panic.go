package fmt

import (
	"fmt"
)

func Panic(a ...interface{}) {
	panic(fmt.Sprint(a...))
}

func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}
