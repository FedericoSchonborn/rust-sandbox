package a

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
	children []any // Element | string
}

func New(href string, children ...any) *Element {
	return &Element{href: href, children: children}
}

func (a *Element) Target(target Target) *Element {
	a.target = target
	return a
}

func (a *Element) Render() error {
	return nil
}
