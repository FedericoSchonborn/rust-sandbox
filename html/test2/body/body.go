package body

import (
	"io"

	"github.com/fdschonborn/sandbox/html/test2"
)

var _ test2.Element = (*Body)(nil)

type Body struct {
	children []test2.Element
}

func New() *Body {
	return new(Body)
}

func (body *Body) Children(elems ...test2.Element) *Body {
	body.children = append(body.children, elems...)
	return body
}

func (*Body) Render(_ io.Writer) error {
	panic("not implemented")
}

type Option func(*Body)
