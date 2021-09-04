package cmp

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Ordering int

const (
	Less    Ordering = -1
	Equal   Ordering = 0
	Greater Ordering = 1
)

func Cmp[T Ordered](l, r T) Ordering {
	switch {
	case l == r:
		return Equal
	case l < r:
		return Less
	case l > r:
		return Greater
	}

	panic("unreachable")
}
