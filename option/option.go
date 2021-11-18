package option

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fdschonborn/go-sandbox/zero"
)

type Option[T any] struct {
	inner *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		inner: &value,
	}
}

// None and Option{} are equivalent.
func None[T any]() Option[T] {
	return Option[T]{
		inner: nil,
	}
}

func (o Option[T]) IsSome() bool {
	return o.inner != nil
}

func (o Option[T]) IsNone() bool {
	return o.inner == nil
}

func (o Option[T]) Get() (value T, ok bool) {
	if o.inner != nil {
		return *o.inner, true
	}

	return zero.Zero[T](), false
}

func (o Option[T]) Unwrap() T {
	if o.inner != nil {
		return *o.inner
	}

	panic("Called Option.Unwrap on None value")
}

// UnwrapUnchecked unwraps the Option without checking if it is Some, this may
// result on unexpected behavior or a panic.
func (o Option[T]) UnwrapUnchecked() T {
	return *o.inner
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.inner != nil {
		return *o.inner
	}

	return def
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.inner != nil {
		return *o.inner
	}

	return fn()
}

// UnwrapOrZero is unwrap_or_default but more gopher-ish.
func (o Option[T]) UnwrapOrZero() T {
	if o.inner != nil {
		return *o.inner
	}

	return zero.Zero[T]()
}

func Map[T, U any](o Option[T], fn func(T) U) Option[U] {
	if o.inner != nil {
		return Some(fn(*o.inner))
	}

	return None[U]()
}

func (o Option[T]) OkOr(err error) (T, error) {
	if o.inner != nil {
		return *o.inner, nil
	}

	return zero.Zero[T](), err
}

func (o Option[T]) OkOrElse(err func() error) (T, error) {
	if o.inner != nil {
		return *o.inner, nil
	}

	return zero.Zero[T](), err()
}

func (o Option[T]) Filter(pred func(T) bool) Option[T] {
	if o.inner != nil && pred(*o.inner) {
		return Some(*o.inner)
	}

	return None[T]()
}

func (o Option[T]) Or(ob Option[T]) Option[T] {
	if o.inner != nil {
		return Some(*o.inner)
	}

	return ob
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.inner != nil {
		return Some(*o.inner)
	}

	return fn()
}

type Zipped[L, R any] struct {
	L L
	R R
}

func Zip[L, R any](l Option[L], r Option[R]) Option[Zipped[L, R]] {
	if l.inner != nil && r.inner != nil {
		return Some(Zipped[L, R]{
			L: *l.inner,
			R: *r.inner,
		})
	}

	return None[Zipped[L, R]]()
}

// TODO: Flatten

func (o Option[T]) Format(f fmt.State, verb rune) {
	if o.inner != nil {
		format := strings.Builder{}
		format.WriteString("Some %")

		if w, ok := f.Width(); ok {
			format.WriteString(strconv.Itoa(w))
		}

		if p, ok := f.Precision(); ok {
			format.WriteByte('.')
			format.WriteString(strconv.Itoa(p))
		}

		if f.Flag('+') {
			format.WriteByte('+')
		}

		if f.Flag('-') {
			format.WriteByte('-')
		}

		if f.Flag('#') {
			format.WriteByte('#')
		}

		if f.Flag(' ') {
			format.WriteByte(' ')
		}

		if f.Flag('0') {
			format.WriteByte('0')
		}

		format.WriteRune(verb)
		fmt.Fprintf(f, format.String(), *o.inner)
		return
	}

	io.WriteString(f, "None")
}
