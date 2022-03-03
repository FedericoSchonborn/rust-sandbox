package task_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fdschonborn/sandbox/go/task"
)

func Example() {
	waitOneSecondAndReturnFive := func() (int, error) {
		time.Sleep(1 * time.Second)
		return 5, nil
	}
	handle := task.New(waitOneSecondAndReturnFive)

	result, err := handle.Join()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	// Output:
	// 5
}

func ExampleJoin() {
	rand.Seed(time.Now().Unix())
	returnSix := func() (int, error) {
		return 6, nil
	}

	handles := make([]*task.Handle[int], 10)
	for i := 0; i < 10; i++ {
		handles[i] = task.New(returnSix)
	}

	array, err := task.Join(handles...)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(array)
	// Output:
	// [6 6 6 6 6 6 6 6 6 6]
}
