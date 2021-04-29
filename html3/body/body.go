package body

import (
	"io"

	"github.com/fdschonborn/x/html3"
)

var _ html3.Element = (*Body)(nil)

type Body struct {
	children []html3.Element
}

func New() *Body {
	return new(Body)
}

func (body *Body) Children(elems ...html3.Element) *Body {
	body.children = append(body.children, elems...)
	return body
}

func (*Body) Render(_ io.Writer) error {
	panic("not implemented")
}

type Option func(*Body)
