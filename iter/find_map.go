package iter

import (
	"github.com/fdschonborn/go-sandbox/option"
)

type FindMapFunc[T, U any] func(T) option.Option[U]

func FindMap[T, U any](iter Iterator[T], f FindMapFunc[T, U]) option.Option[U] {
	// The func is a fugly workaround because FilterMapFunc != FindMapFunc
	return FilterMap(iter, func(item T) option.Option[U] { return f(item) }).Next()
}
