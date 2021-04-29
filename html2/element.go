package html2

import "io"

type Element interface {
	Render(io.Writer) error
}
