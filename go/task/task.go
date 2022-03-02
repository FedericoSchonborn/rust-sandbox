package task

import (
	"sync"
)

type Handle[T any] struct {
	done   chan bool
	result T
}

func Join[T any](handles ...*Handle[T]) []T {
	var (
		wg      sync.WaitGroup
		results = make([]T, len(handles))
	)

	for n, handle := range handles {
		wg.Add(1)

		go func(handle *Handle[T], results []T, n int) {
			results[n] = handle.Join()
			wg.Done()
		}(handle, results, n)
	}

	wg.Wait()
	return results
}

func New[T any](fn func() T) *Handle[T] {
	handle := &Handle[T]{
		done: make(chan bool),
	}

	go func() {
		handle.result = fn()
		handle.done <- true
	}()
	return handle
}

func (h *Handle[T]) Join() T {
	<-h.done
	return h.result
}
