package types

import "math"

type Int int

func (i Int) Pow(y Int) Int {
	return Int(math.Pow(float64(i), float64(y)))
}
