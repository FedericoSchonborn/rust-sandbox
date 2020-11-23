package types

import "math"

type Float64 float64

func (f Float64) Pow(y Float64) Float64 {
	return Float64(math.Pow(float64(f), float64(y)))
}
