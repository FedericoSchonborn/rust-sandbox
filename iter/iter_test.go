package iter_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/iter"
)

func ExampleFind() {
	a := []int{1, 2, 3}

	equals := func(value int) func(int) bool {
		return func(item int) bool {
			return item == value
		}
	}

	fmt.Println(iter.Find(iter.SliceIterator(a), equals(2)))
	fmt.Println(iter.Find(iter.SliceIterator(a), equals(5)))
	// Output:
	// Some(2)
	// None
}

func ExampleFold() {
	a := []int{1, 2, 3}
	sum := iter.Fold(iter.SliceIterator(a), 0, func(acc int, item int) int {
		return acc + item
	})

	fmt.Println(sum)
	// Output:
	// 6
}

/* NOTE: Broken
func ExampleMap() {
	a := []int{1, 2, 3}
	iter := iter.Map(iter.SliceIterator(a), func(item int) int {
		return 2 * item
	})

	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(2)
	// Some(4)
	// Some(6)
	// None
}
*/
