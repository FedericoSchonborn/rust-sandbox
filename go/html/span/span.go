package span

type Element struct {
	children []any // Element | string
}

func New(children ...any /* Element | string */) *Element {
	return &Element{children}
}

func (span *Element) Render() error {
	return nil
}
