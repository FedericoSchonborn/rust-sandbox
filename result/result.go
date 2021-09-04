package result

import "github.com/fdschonborn/go-sandbox/zero"

type Result[T any, E error] struct {
	ok   bool
	data interface{}
}

func Ok[T any, E error](value T) Result[T, E] {
	return Result[T, E]{
		ok:   true,
		data: value,
	}
}

func Err[T any, E error](err error) Result[T, E] {
	return Result[T, E]{
		ok:   false,
		data: err,
	}
}

func (r Result[T, E]) IsOk() bool {
	return r.ok
}

func (r Result[T, E]) IsErr() bool {
	return !r.ok
}

func (r Result[T, E]) Unpack() (T, error) {
	if r.ok {
		return r.data.(T), nil
	}

	return zero.Zero[T](), r.data.(error)
}
