package html

import "io"

type Div struct {
	Content []Element
}

var _ Element = (*Div)(nil)

func (div *Div) Render(w io.Writer) error {
	if _, err := io.WriteString(w, "<div>"); err != nil {
		return err
	}

	for _, elem := range div.Content {
		if err := elem.Render(w); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, "</div>"); err != nil {
		return err
	}

	return nil
}
