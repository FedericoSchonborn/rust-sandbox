package main

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/tuple"
)

// NOTE: This crashes the gopher.
func main() {
	inner := map[string]int{
		"banana": 42,
		"apple":  38,
		"orange": 96,
	}

	iter.IntoSlice[tuple.Tuple2[string, int]](
		iter.ForEach(
			iter.NewMapIterator(inner), func(tuple tuple.Tuple2[string, int]) {
				fmt.Println(tuple.A, tuple.B)
			},
		),
	)
}
