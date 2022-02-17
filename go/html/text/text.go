package text

type Element struct {
	s string
}

func New(s string) *Element {
	return &Element{s: s}
}

func (text *Element) Render() error {
	return nil
}
