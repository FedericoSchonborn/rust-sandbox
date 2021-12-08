package result

import (
	"fmt"

	xfmt "github.com/fdschonborn/go-sandbox/fmt"
)

type Result[T any, E error] struct {
	ok    bool
	inner interface{}
}

func Ok[T any, E error](value T) Result[T, E] {
	return Result[T, E]{
		ok:    true,
		inner: value,
	}
}

// Result{} and Err(nil) are equivalent.
func Err[T any, E error](err error) Result[T, E] {
	return Result[T, E]{
		ok:    false,
		inner: err,
	}
}

func (r Result[T, E]) IsOk() bool {
	return r.ok
}

func (r Result[T, E]) IsErr() bool {
	return !r.ok
}

func Map[T any, E error, U any](r Result[T, E], op func(T) U) Result[U, E] {
	if !r.ok {
		return Err[U, E](r.inner.(E))
	}

	return Ok[U, E](op(r.inner.(T)))
}

func (r Result[T, E]) Format(f fmt.State, verb rune) {
	var kind string
	if r.ok {
		kind = "Ok "
	} else {
		kind = "Err "
	}

	fmt.Fprintf(f, kind+xfmt.Format(f, verb), r.inner)
}
