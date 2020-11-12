package is

import (
	"math"

	"github.com/fdschonborn/x/go/typeof"
)

func Even(i interface{}) bool {
	switch v := i.(type) {
	case int:
		return v%2 == 0
	case uint:
		return v%2 == 0
	case int8:
		return v%2 == 0
	case uint8:
		return v%2 == 0
	case int16:
		return v%2 == 0
	case uint16:
		return v%2 == 0
	case int32:
		return v%2 == 0
	case uint32:
		return v%2 == 0
	case int64:
		return v%2 == 0
	case uint64:
		return v%2 == 0
	case float32:
		return math.Mod(float64(v), 2) == 0
	case float64:
		return math.Mod(v, 2) == 0
	default:
		panic("Expected numerical value, got " + typeof.TypeOf(i) + " instead")
	}
}

func Odd(i interface{}) bool {
	return !Even(i)
}
