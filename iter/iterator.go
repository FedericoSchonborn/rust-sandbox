package iter

import "github.com/fdschonborn/go-sandbox/option"

type Iterator[Item any] interface {
	Next() option.Option[Item]
}

func ForEach[T any, I Iterator[T]](iter I, f func(T)) {
	for value := iter.Next(); value != option.None[T](); value = iter.Next() {
		f(value.UnwrapUnchecked())
	}
}
