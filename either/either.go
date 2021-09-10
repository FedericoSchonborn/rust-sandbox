package either

import "github.com/fdschonborn/go-sandbox/zero"

type Either[L, R any] struct {
	left bool
	data interface{}
}

func Left[L, R any](value L) Either[L, R] {
	return Either[L, R]{
		left: true,
		data: value,
	}
}

func Right[L, R any](value R) Either[L, R] {
	return Either[L, R]{
		left: false,
		data: value,
	}
}

func (e Either[L, R]) IsLeft() bool {
	return e.left
}

func (e Either[L, R]) IsRight() bool {
	return !e.left
}

func (e Either[L, R]) Unpack() (L, R) {
	if e.left {
		return e.data.(L), zero.Zero[R]()
	}

	return zero.Zero[L](), e.data.(R)
}
