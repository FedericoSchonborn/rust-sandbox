package iter

import "github.com/fdschonborn/go-sandbox/option"

type forEach[T any, Iter Iterator[T]] struct {
	iter     Iter
	callback func(T)
}

func ForEach[T any, Iter Iterator[T]](iter Iter, f func(T)) Iterator[T] {
	return &forEach[T, Iter]{
		iter:     iter,
		callback: f,
	}
}

func (fei *forEach[T, Iter]) Next() option.Option[T] {
	// TODO: Replace this crap with something more appropriate.
	return option.Map(fei.iter.Next(), func(value T) T {
		fei.callback(value)
		return value
	})
}
