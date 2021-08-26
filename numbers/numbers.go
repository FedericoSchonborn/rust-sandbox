package numbers

import (
	"github.com/fdschonborn/go-sandbox/floats"
	"github.com/fdschonborn/go-sandbox/ints"
)

type Number interface {
	ints.Int | floats.Float
}
