package zero_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/zero"
)

type Banana string

func (b Banana) Default() Banana {
	return Banana("banana")
}

func ExampleDefault() {
	fmt.Println(zero.Default[Banana]())
	fmt.Println(zero.Default[int]())
	// Output:
	// banana
	// 0
}
