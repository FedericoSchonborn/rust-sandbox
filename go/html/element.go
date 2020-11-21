package html

import (
	"io"
)

type Element interface {
	Render(io.Writer) error
}
