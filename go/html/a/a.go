package a

import "github.com/fdschonborn/sandbox/go/html"

type Target string

const (
	Blank  Target = "_blank"
	Self   Target = "_self"
	Parent Target = "_parent"
	Top    Target = "_top"
)

type Element struct {
	href     string
	target   Target
	children []html.Element
}

func New(href string, children ...html.Element) *Element {
	return &Element{href: href, children: children}
}

func (a *Element) Target(target Target) *Element {
	a.target = target
	return a
}

func (a *Element) Render() error {
	return nil
}
