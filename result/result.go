package result

import "github.com/fdschonborn/go-sandbox/zero"

type tag uint8

const (
	tagOk tag = iota + 1
	tagErr
)

type Result[T any, E error] struct {
	tag  tag
	data interface{}
}

func Ok[T any, E error](value T) Result[T, E] {
	return Result[T, E]{
		tag:  tagOk,
		data: value,
	}
}

func Err[T any, E error](err error) Result[T, E] {
	return Result[T, E]{
		tag:  tagErr,
		data: err,
	}
}

func (r Result[T, E]) IsOk() bool {
	return r.tag == tagOk
}

func (r Result[T, E]) IsErr() bool {
	return r.tag == tagErr
}

func (r Result[T, E]) Unpack() (T, error) {
	switch r.tag {
	case tagOk:
		return r.data.(T), nil
	case tagErr:
		return zero.Zero[T](), r.data.(error)
	default:
		panic("Result contains invalid tag")
	}
}
