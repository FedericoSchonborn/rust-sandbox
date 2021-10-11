package promise_test

import (
	"fmt"
	"os"

	"github.com/fdschonborn/go-sandbox/promise"
)

func Example() {
	promise.Resolve("Banana").Then(func(value string) {
		fmt.Println(value)
	}).Catch(func (err error) {
		fmt.Fprintln(os.Stderr, err)
	})

	// Output:
	// Banana
}
