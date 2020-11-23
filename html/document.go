package html

import (
	"io"
)

type Document struct {
	head *HeadElement
	body *BodyElement
}

func New(head *HeadElement, body *BodyElement) *Document {
	return &Document{
		head: head,
		body: body,
	}
}

func (doc *Document) Head() *HeadElement {
	return doc.head
}

func (doc *Document) Body() *BodyElement {
	return doc.body
}

func (doc *Document) Render(w io.Writer) error {
	headctx := NewContext(doc.head, nil, w)
	if err := doc.head.Render(headctx); err != nil {
		return err
	}

	bodyctx := NewContext(doc.body, nil, w)
	if err := doc.body.Render(bodyctx); err != nil {
		return err
	}

	return nil
}
