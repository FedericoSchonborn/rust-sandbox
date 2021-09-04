package iter_test

import (
	"fmt"
	"strconv"

	"github.com/fdschonborn/go-sandbox/iter"
	"github.com/fdschonborn/go-sandbox/option"
)

func ExampleFind() {
	a := []int{1, 2, 3}

	equals := func(value int) func(int) bool {
		return func(item int) bool {
			return item == value
		}
	}

	fmt.Println(iter.Find(iter.FromSlice(a), equals(2)))
	fmt.Println(iter.Find(iter.FromSlice(a), equals(5)))
	// Output:
	// Some(2)
	// None
}

func ExampleFold() {
	a := []int{1, 2, 3}
	sum := iter.Fold(iter.FromSlice(a), 0, func(acc int, item int) int {
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

func ExampleFilter() {
	a := []int{0, 1, 2}
	iter := iter.Filter(iter.FromSlice(a), func(item int) bool {
		return item > 0
	})

	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(1)
	// Some(2)
	// None
}

func ExampleFilterMap() {
	a := []string{"1", "two", "NaN", "four", "5"}
	iter := iter.FilterMap(iter.FromSlice(a), func(item string) option.Option[int] {
		n, err := strconv.Atoi(item)
		if err != nil {
			return option.None[int]()
		}

		return option.Some(n)
	})

	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// Some(1)
	// Some(5)
	// None
}

func ExampleFindMap() {
	a := []string{"lol", "NaN", "2", "5"}
	firstNumber := iter.FindMap(iter.FromSlice(a), func(item string) option.Option[int] {
		n, err := strconv.Atoi(item)
		if err != nil {
			return option.None[int]()
		}

		return option.Some(n)
	})

	fmt.Println(firstNumber)
	// Output:
	// Some(2)
}