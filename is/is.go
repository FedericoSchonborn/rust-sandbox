package is

import (
	"math"

	"github.com/fdschonborn/x/fmt"
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
		fmt.Panicf("Expected numerical value, got %T", v)
		return false
	}
}

func Odd(i interface{}) bool {
	return !Even(i)
}
