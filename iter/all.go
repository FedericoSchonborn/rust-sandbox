package iter

type AllFunc[T any] func(T) bool

func All[T any](iter Iterator[T], f AllFunc[T]) bool {
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			break
		}

		if !f(item) {
			return false
		}
	}

	return true
}
