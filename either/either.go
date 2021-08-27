package either

import "github.com/fdschonborn/go-sandbox/zero"

type tag uint8

const (
	tagLeft tag = iota + 1
	tagRight
)

type Either[L, R any] struct {
	tag  tag
	data interface{}
}

func Left[L, R any](value L) Either[L, R] {
	return Either[L, R]{
		tag:  tagLeft,
		data: value,
	}
}

func Right[L, R any](value R) Either[L, R] {
	return Either[L, R]{
		tag:  tagRight,
		data: value,
	}
}

func (e Either[L, R]) IsLeft() bool {
	return e.tag == tagLeft
}

func (e Either[L, R]) IsRight() bool {
	return e.tag == tagRight
}

func (e Either[L, R]) Unpack() (L, R) {
	switch e.tag {
	case tagLeft:
		return e.data.(L), zero.Zero[R]()
	case tagRight:
		return zero.Zero[L](), e.data.(R)
	default:
		panic("Either contains invalid tag")
	}
}
