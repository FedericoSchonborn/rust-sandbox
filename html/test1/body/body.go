package body

import (
	"io"

	"github.com/fdschonborn/sandbox/html/test1"
)

var _ test1.Element = (*Body)(nil)

type Body struct {
	children []test1.Element
}

func New(opts ...Option) *Body {
	body := new(Body)
	for _, opt := range opts {
		opt(body)
	}

	return body
}

func (*Body) Render(_ io.Writer) error {
	panic("not implemented")
}

type Option func(*Body)

func Children(elems ...test1.Element) Option {
	return func(body *Body) {
		body.children = append(body.children, elems...)
	}
}
