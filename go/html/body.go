package html

import (
	"io"
)

type BodyElement struct {
	content []BodyChild
}

func Body(content ...BodyChild) *BodyElement {
	return &BodyElement{
		content: content,
	}
}

func (body *BodyElement) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "<body>"); err != nil {
		return err
	}

	for _, elem := range body.content {
		if err := elem.Render(nil, w); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, "</body>"); err != nil {
		return err
	}

	return nil
}
