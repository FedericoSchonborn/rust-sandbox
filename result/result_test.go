package result_test

import (
	"fmt"

	"github.com/fdschonborn/sandbox/result"
	"github.com/fdschonborn/sandbox/result/strconv"
)

func ExampleMap() {
	value := result.Map(strconv.Atoi("2"), func(i int) int {
		return i * 2
	})
	fmt.Println(value)

	// Output:
	// Ok 4
}
