package odd

import (
	"github.com/fdschonborn/go-sandbox/ints"
	"github.com/fdschonborn/go-sandbox/is/even"
)

func IsOdd[T ints.Int](value T) bool {
	return !even.IsEven(value)
}
