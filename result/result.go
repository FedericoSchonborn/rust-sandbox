package result

type Result[T any] struct {
	ok    bool
	inner interface{}
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		ok:    true,
		inner: value,
	}
}

// Result{} and Err(nil) are equivalent.
func Err[T any](err error) Result[T] {
	return Result[T]{
		ok:    false,
		inner: err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.ok
}

func (r Result[T]) IsErr() bool {
	return !r.ok
}

// unsafe
func (r Result[T]) UnwrapUnchecked() T {
	return r.inner.(T)
}

func Map[T, U any](r Result[T], op func(T) U) Result[U] {
	if !r.ok {
		return Err[U](r.inner.(error))
	}

	return Ok(op(r.inner.(T)))
}
