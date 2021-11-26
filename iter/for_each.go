package iter

type ForEachFunc[T any] func(T)

func ForEach[T any](iter Iterator[T], f ForEachFunc[T]) {
	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			break
		}

		f(item)
	}
}
