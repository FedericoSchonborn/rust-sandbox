package p

import "io"

type P struct {
	text string
}

func New(opts ...Option) *P {
	p := new(P)
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (*P) Render(_ io.Writer) error {
	panic("not implemented")
}

type Option func(*P)

func Text(text string) Option {
	return func(p *P) {
		p.text = text
	}
}
