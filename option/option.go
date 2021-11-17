package option

import (
	"fmt"
	"io"

	"github.com/fdschonborn/go-sandbox/zero"
)

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		some:  true,
		value: value,
	}
}

// None and Option{} are equivalent.
func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Unwrap() T {
	if o.some {
		return o.value
	}

	panic("Called Option.Unwrap on None value")
}

// UnwrapUnchecked unwraps the Option without checking if it is Some, this may
// result on unexpected behavior or a panic.
func (o Option[T]) UnwrapUnchecked() T {
	return o.value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.value
	}

	return def
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.some {
		return o.value
	}

	return fn()
}

// UnwrapOrZero is unwrap_or_default but more gopher-ish.
func (o Option[T]) UnwrapOrZero() T {
	if o.some {
		return o.value
	}

	return zero.Zero[T]()
}

func Map[T, U any](o Option[T], fn func(T) U) Option[U] {
	if o.some {
		return Some(fn(o.value))
	}

	return None[U]()
}

func (o Option[T]) OkOr(err error) (T, error) {
	if o.some {
		return o.value, nil
	}

	return zero.Zero[T](), err
}

func (o Option[T]) OkOrElse(err func() error) (T, error) {
	if o.some {
		return o.value, nil
	}

	return zero.Zero[T](), err()
}

func (o Option[T]) Filter(pred func(T) bool) Option[T] {
	if o.some && pred(o.value) {
		return Some(o.value)
	}

	return None[T]()
}

func (o Option[T]) Or(ob Option[T]) Option[T] {
	if o.some {
		return o
	}

	return ob
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.some {
		return o
	}

	return fn()
}

type Zipped[L, R any] struct {
	L L
	R R
}

func Zip[L, R any](lhs Option[L], rhs Option[R]) Option[Zipped[L, R]] {
	if lhs.some && rhs.some {
		return Some(Zipped[L, R]{
			L: lhs.value,
			R: rhs.value,
		})
	}

	return None[Zipped[L, R]]()
}

// TODO: Flatten

func (o Option[T]) Format(f fmt.State, verb rune) {
	if o.some {
		io.WriteString(f, "Some("+fmt.Sprintf("%#v", o.value)+")")
		return
	}

	io.WriteString(f, "None")
}
