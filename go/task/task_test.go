package task_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fdschonborn/sandbox/go/task"
)

func Example() {
	waitOneSecondAndReturnFive := func() int {
		time.Sleep(1 * time.Second)
		return 5
	}

	handle := task.New(waitOneSecondAndReturnFive)
	result := handle.Join()
	fmt.Println(result)

	// Output:
	// 5
}

func ExampleJoin() {
	rand.Seed(time.Now().Unix())
	returnSix := func() int {
		return 6
	}

	handles := make([]*task.Handle[int], 10)
	for i := 0; i < 10; i++ {
		handles[i] = task.New(returnSix)
	}

	results := task.Join(handles...)
	fmt.Println(results)

	// Output:
	// [6 6 6 6 6 6 6 6 6 6]
}
