package test3

type BodyOption interface {
	bodyOption(*BodyElement) error
}

type BodyElement struct {
	children []Element
}

func Body(opts ...BodyOption) *BodyElement {
	body := &BodyElement{}

	for _, opt := range opts {
		if err := opt.bodyOption(body); err != nil {
			panic(err)
		}
	}

	return body
}

func (b *BodyElement) element() {}
