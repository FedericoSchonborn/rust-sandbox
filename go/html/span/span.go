package span

import "github.com/fdschonborn/sandbox/go/html"

type Element struct {
	children []html.Element
}

func New(children ...html.Element) *Element {
	return &Element{children}
}

func (span *Element) Render() error {
	return nil
}
