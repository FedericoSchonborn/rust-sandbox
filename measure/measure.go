package measure

import "github.com/fdschonborn/go-sandbox/numbers"

type Meter[T numbers.Number] T

func M[T numbers.Number](value T) Meter[T] {
	return Meter[T](value)
}

func (m Meter[T]) Km() Kilometer[T] {
	return Kilometer[T](T(m) / T(1000))
}

type Kilometer[T numbers.Number] T

func Km[T numbers.Number](value T) Kilometer[T] {
	return Kilometer[T](value)
}

func (km Kilometer[T]) M() Meter[T] {
	return Meter[T](T(km) * T(1000))
}
