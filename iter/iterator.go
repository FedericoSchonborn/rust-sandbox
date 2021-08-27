package iter

import "github.com/fdschonborn/go-sandbox/option"

type Iterator[Item any] interface {
	Next() option.Option[Item]
}
