package iter

type AnyFunc[T any] func(T) bool

func Any[T any](iter Iterator[T], f AnyFunc[T]) bool {
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			break
		}

		if f(item) {
			return true
		}
	}

	return false
}
