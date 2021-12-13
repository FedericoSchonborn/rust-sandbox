package task

import "sync"

type Task[T any] struct {
	result T
	done   chan struct{}
}

func New[T any](fn func() T) *Task[T] {
	t := &Task[T]{
		done: make(chan struct{}),
	}

	go func() {
		t.result = fn()
		t.done <- struct{}{}
	}()

	return t
}

func (t *Task[T]) Await() T {
	<-t.done
	return t.result
}

func All[T any](ts ...*Task[T]) []T {
	rs := make([]T, len(ts))
	var wg sync.WaitGroup
	for i, t := range ts {
		go func(i int, t *Task[T]) {
			defer wg.Done()
			wg.Add(1)

			rs[i] = t.Await()
		}(i, t)
	}

	wg.Wait()
	return rs
}
