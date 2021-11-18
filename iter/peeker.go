package iter

import "github.com/fdschonborn/go-sandbox/option"

type Peeker[T any] interface {
	Iterator[T]

	Peek() option.Option[T]
}
