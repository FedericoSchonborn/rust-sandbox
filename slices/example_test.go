package slices_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/slices"
)

func ExampleSwap() {
	v := []string{"a", "b", "c", "d"}
	slices.Swap(v, 1, 3)

	fmt.Println(v)
	// Output:
	// [a d c b]
}

func ExampleRepeat() {
	fmt.Println(slices.Repeat([]int{1, 2}, 3))
	// Output:
	// [1 2 1 2 1 2]
}
