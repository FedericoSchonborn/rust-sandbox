package task

import (
	"sync"
	"sync/atomic"
)

type Handle[T any] struct {
	done  chan bool
	value T
	err   error
}

func Join[T any](handles ...*Handle[T]) ([]T, error) {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		values []T
		err    error
		stop   int32
	)

	for _, handle := range handles {
		wg.Add(1)

		go func(handle *Handle[T]) {
			defer wg.Done()

			if atomic.LoadInt32(&stop) == 1 {
				return
			}

			value, _err := handle.Join()
			if _err != nil {
				mu.Lock()
				err = _err
				mu.Unlock()

				atomic.StoreInt32(&stop, 1)
				return
			}

			mu.Lock()
			values = append(values, value)
			mu.Unlock()
		}(handle)
	}

	wg.Wait()
	return values, err
}

func New[T any](fn func() (T, error)) *Handle[T] {
	handle := &Handle[T]{
		done: make(chan bool),
	}

	go func() {
		handle.value, handle.err = fn()
		handle.done <- true
	}()
	return handle
}

func (h *Handle[T]) Join() (T, error) {
	<-h.done
	return h.value, h.err
}
