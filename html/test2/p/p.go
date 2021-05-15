package p

import "io"

type P struct {
	text string
}

func New() *P {
	return new(P)
}

func (p *P) Text(text string) *P {
	p.text = text
	return p
}

func (*P) Render(_ io.Writer) error {
	panic("not implemented")
}
