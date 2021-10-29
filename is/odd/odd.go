package odd

import (
	"github.com/fdschonborn/go-sandbox/constraints"
	"github.com/fdschonborn/go-sandbox/is/even"
)

func IsOdd[T constraints.Int](value T) bool {
	return !even.IsEven(value)
}
