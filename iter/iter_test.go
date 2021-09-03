package iter_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/tuple"
)

func Example() {
	inner := map[string]int{
		"banana": 42,
		"apple":  38,
		"orange": 96,
	}

	iter.ForEach(
		iter.MapIterator(inner), func(tuple tuple.Tuple2[string, int]) {
			key, value := tuple.Unpack()

			fmt.Println(key, value)
		},
	)
}
