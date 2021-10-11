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
	defer close(t.ch)

	return <-t.ch
}

type TryTask[T any] struct {
	ch  chan T
	err error
}

func Try[T any](fn func() (T, error)) TryTask[T] {
	var err error

	ch := make(chan T)
	go func() {
		value, inerr := fn()
		if inerr != nil {
			err = inerr
			return
		}

		ch <- value
	}()

	return TryTask[T]{
		ch:  ch,
		err: err,
	}
}

func (tt TryTask[T]) Await() (T, error) {
	defer close(tt.ch)

	return <-tt.ch, tt.err
}
