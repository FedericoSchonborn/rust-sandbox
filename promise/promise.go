package promise

type ThenFunc[T any] func(T)
type CatchFunc func(error)

type Promise[T any] interface {
	Then(ThenFunc[T]) Promise[T]
	Catch(CatchFunc) Promise[T]
}

type promise[T any] struct {
	out T
	err error
}

func Resolve[T any](value T) Promise[T] {
	return promise[T]{
		out: value,
	}
}

func Reject[T any](err error) Promise[T] {
	return promise[T]{
		err: err,
	}
}

func (p promise[T]) Then(fn ThenFunc[T]) Promise[T] {
	panic("todo")
}

func (p promise[T]) Catch(fn CatchFunc) Promise[T] {
	panic("todo")
}
