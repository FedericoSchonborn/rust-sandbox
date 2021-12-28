package html

type SpanElement struct {
	children []any // Element | string
}

func Span(children ...any /* Element | string */) *SpanElement {
	return &SpanElement{children}
}

func (span *SpanElement) Render() error {
	return nil
}
