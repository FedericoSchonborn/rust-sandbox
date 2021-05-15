package document

import (
	"io"

	"github.com/fdschonborn/x/html/test1"
	"github.com/fdschonborn/x/html/test1/body"
)

var _ test1.Element = (*Document)(nil)

type Document struct {
	body *body.Body
}

func New(opts ...Option) *Document {
	doc := new(Document)
	for _, opt := range opts {
		opt(doc)
	}

	return doc
}

func (*Document) Render(_ io.Writer) error {
	panic("not implemented")
}

type Option func(*Document)

func Body(body *body.Body) Option {
	return func(d *Document) {
		d.body = body
	}
}
