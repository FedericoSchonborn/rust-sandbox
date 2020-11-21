package html

import "io"

var _ BodyChild = (*DivElement)(nil)

type DivElement struct {
	content []BodyChild
	class   string
}

func Div(content ...BodyChild) *DivElement {
	return &DivElement{
		content: content,
	}
}

func (div *DivElement) Class(class string) *DivElement {
	div.class = class
	return div
}

func (div *DivElement) Render(body *BodyElement, w io.Writer) error {
	openTag := "<div"
	if div.class != "" {
		openTag += " class=\"" + div.class + "\""
	}
	openTag += ">"

	if _, err := io.WriteString(w, openTag); err != nil {
		return err
	}

	for _, elem := range div.content {
		if err := elem.Render(body, w); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, "</div>"); err != nil {
		return err
	}

	return nil
}
