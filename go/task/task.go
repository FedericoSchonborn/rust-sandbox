package task

import "sync"

type Handle[T any] struct {
	done   chan bool
	result T
}

func Join[T any](handles ...*Handle[T]) []T {
	var wg sync.WaitGroup
	var m sync.Mutex
	results := make([]T, len(handles))

	for _, handle := range handles {
		wg.Add(1)

		go func(handle *Handle[T]) {
			result := handle.Join()

			m.Lock()
			results = append(results, result)
			m.Unlock()

			wg.Done()
		}(handle)
	}

	wg.Wait()
	return results
}

func Spawn[T any](fn func() T) *Handle[T] {
	done := make(chan bool)
	var result T
	go func() {
		result = fn()
		done <- true
	}()

	return &Handle[T]{
		done:   done,
		result: result,
	}
}

func (h *Handle[T]) Join() T {
	<-h.done
	return h.result
}
