package html3

import "io"

type Element interface {
	Render(io.Writer) error
}
