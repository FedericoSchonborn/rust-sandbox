package body

import (
	"io"

	"github.com/fdschonborn/x/html2"
)

var _ html2.Element = (*Body)(nil)

type Body struct {
	children []html2.Element
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

func Children(elems ...html2.Element) Option {
	return func(body *Body) {
		body.children = append(body.children, elems...)
	}
}
