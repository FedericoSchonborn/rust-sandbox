package task_test

import (
	"fmt"
	"time"

	"github.com/fdschonborn/go-sandbox/task"
	"github.com/fdschonborn/go-sandbox/tuple"
)

func Example() {
	getOne := func() int {
		return 1
	}

	task := task.New(getOne)
	fmt.Println(task.Await())

	// Output:
	// 1
}

func ExampleSleep() {
	sleep := func() struct{} {
		time.Sleep(5 * time.Second)
		return struct{}{}
	}

	task := task.New(sleep)
	fmt.Println(task.Await())

	// Output:
	// {}
}

func ExampleTryTask() {
	withTuple := func() (tuple.Tuple3[string, int, bool], error) {
		return tuple.New3("Hello, world!", 42, true), nil
	}

	task := task.Try(withTuple)
	fmt.Println(task.Await())

	// Output:
	// {Hello, world! 42 true} <nil>
}
