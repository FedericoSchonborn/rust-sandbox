package html

type PElement struct {
	children []any // Element | string
}

func P(children ...any /* Element | string */) *PElement {
	return &PElement{children}
}

func (p *PElement) Render() error {
	return nil
}
