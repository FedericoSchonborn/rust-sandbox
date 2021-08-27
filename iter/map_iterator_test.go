package iter_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/tuple"
)

func ExampleMapIterator() {
	inner := map[string]int{
		"banana": 42,
		"apple":  38,
		"orange": 96,
	}

	iter.ForEach(iter.NewMapIterator(inner), func(tuple tuple.Tuple2[string, int]) {
		fmt.Println(tuple.A, tuple.B)
	})
	// Output:
	// TEST
}
