package option

import (
	"github.com/fdschonborn/go-sandbox/tuple"
	"github.com/fdschonborn/go-sandbox/zero"
)

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None and Option{} are equivalent.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Unwrap() T {
	if o.value == nil {
		panic("Called Option.Unwrap on None value")
	}

	return *o.value
}

// UnwrapUnchecked unwraps the Option without checking if it is Some, this may
// result on unexpected behavior or a panic.
func (o Option[T]) UnwrapUnchecked() T {
	return *o.value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.value == nil {
		return def
	}

	return *o.value
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.value == nil {
		return fn()
	}

	return *o.value
}

// UnwrapOrZero is unwrap_or_default but more gopher-ish.
func (o Option[T]) UnwrapOrZero() T {
	if o.value == nil {
		return zero.Zero[T]()
	}

	return *o.value
}

func Map[T, U any](o Option[T], fn func(T) U) Option[U] {
	if o.value == nil {
		return None[U]()
	}

	return Some(fn(*o.value))

}

func (o Option[T]) OkOr(err error) (T, error) {
	if o.value == nil {
		return zero.Zero[T](), err
	}

	return *o.value, nil
}

func (o Option[T]) OkOrElse(fn func() error) (T, error) {
	if o.value == nil {
		return zero.Zero[T](), fn()
	}

	return *o.value, nil
}

func (o Option[T]) Filter(pred func(T) bool) Option[T] {
	if o.value == nil {
		return None[T]()
	}

	if !pred(*o.value) {
		return None[T]()
	}

	return Some(*o.value)
}

func (o Option[T]) Or(ob Option[T]) Option[T] {
	if o.value == nil {
		return ob
	}

	return o
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.value == nil {
		return fn()
	}

	return o
}

func Zip[A, B any](a Option[A], b Option[B]) Option[tuple.Tuple2[A, B]] {
	if a.value != nil && b.value != nil {
		return Some(tuple.Tuple2[A, B]{
			A: *a.value,
			B: *b.value,
		})
	}

	return None[tuple.Tuple2[A, B]]()
}

func Flatten[T any](o Option[Option[T]]) Option[T] {
	if o.value == nil || *o.value == None[T]() {
		return None[T]()
	}

	return *o.value
}
