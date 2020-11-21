package html

import (
	"io"
)

type HeadElement struct{}

func Head() *HeadElement {
	return &HeadElement{}
}

func (head *HeadElement) Render(w io.Writer) error {
	return nil
}
