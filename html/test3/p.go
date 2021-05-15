package test3

type POption interface {
	pOption(*PElement) error
}

type PElement struct {
	text string
}

func P(opts ...POption) *PElement {
	p := &PElement{}

	for _, opt := range opts {
		if err := opt.pOption(p); err != nil {
			panic(err)
		}
	}

	return p
}

func (p *PElement) element() {}
