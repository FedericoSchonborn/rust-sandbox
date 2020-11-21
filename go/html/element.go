package html

import (
	"io"
)

type BodyChild interface {
	Render(*BodyElement, io.Writer) error
}

type HeadChild interface {
	Render(*HeadElement, io.Writer) error
}
