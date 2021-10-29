package even

import "github.com/fdschonborn/go-sandbox/constraints"

func IsEven[T constraints.Int](value T) bool {
	return value%2 == 0
}
