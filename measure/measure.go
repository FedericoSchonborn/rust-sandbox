package measure

import "github.com/fdschonborn/go-sandbox/constraints"

type Meter[T constraints.Number] T

func M[T constraints.Number](value T) Meter[T] {
	return Meter[T](value)
}

func (m Meter[T]) Km() Kilometer[T] {
	return Kilometer[T](T(m) / T(1000))
}

type Kilometer[T constraints.Number] T

func Km[T constraints.Number](value T) Kilometer[T] {
	return Kilometer[T](value)
}

func (km Kilometer[T]) M() Meter[T] {
	return Meter[T](T(km) * T(1000))
}
