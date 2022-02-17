package p

import "github.com/fdschonborn/sandbox/go/html"

type Element struct {
	children []html.Element
}

func New(children ...html.Element) *Element {
	return &Element{children}
}

func (p *Element) Render() error {
	return nil
}
