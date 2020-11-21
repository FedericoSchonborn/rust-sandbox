package html

import (
	"io"
)

type HTML struct {
	head *HeadElement
	body *BodyElement
}

func New(head *HeadElement, body *BodyElement) *HTML {
	return &HTML{
		head: head,
		body: body,
	}
}

func (doc *HTML) Head() *HeadElement {
	return doc.head
}

func (doc *HTML) Body() *BodyElement {
	return doc.body
}

func (doc *HTML) Render(w io.Writer) error {
	if err := doc.head.Render(w); err != nil {
		return err
	}

	if err := doc.body.Render(w); err != nil {
		return err
	}

	return nil
}
