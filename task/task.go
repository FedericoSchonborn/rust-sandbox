package task


type Task[T any] struct {
	ch chan T
}

func New[T any](fn func() T) Task[T] {
	ch := make(chan T)
	go func() {
		ch <- fn()
	}()

	return Task[T]{
		ch: ch,
	}
}

func (t Task[T]) Await() T {
	return <-t.ch
}
