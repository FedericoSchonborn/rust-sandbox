package document

import (
	"io"

	"github.com/fdschonborn/x/html/test2"
	"github.com/fdschonborn/x/html/test2/body"
)

var _ test2.Element = (*Document)(nil)

type Document struct {
	body *body.Body
}

func New() *Document {
	return new(Document)
}

func (doc *Document) Body(body *body.Body) *Document {
	doc.body = body
	return doc
}

func (*Document) Render(_ io.Writer) error {
	panic("not implemented")
}
