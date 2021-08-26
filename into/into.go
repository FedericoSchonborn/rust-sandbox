package into

type Interface[T any] interface {
	Into() T
}

func Into[T any, V Interface[T]](value V) T {
	return value.Into()
}
