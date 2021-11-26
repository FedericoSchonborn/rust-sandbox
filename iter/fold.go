package iter

type FoldFunc[T, A any] func(A, T) A

func Fold[T, A any](iter Iterator[T], init A, f FoldFunc[T, A]) A {
	acc := init

	for {
		item, ok := iter.Next().UnwrapOrZero()
		if !ok {
			return acc
		}

		acc = f(acc, item)
	}
}
