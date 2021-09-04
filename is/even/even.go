package even

import "github.com/fdschonborn/go-sandbox/ints"

func IsEven[T ints.Int](value T) bool {
	return value%2 == 0
}
