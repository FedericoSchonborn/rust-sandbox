package test1

import "io"

type Element interface {
	Render(io.Writer) error
}
