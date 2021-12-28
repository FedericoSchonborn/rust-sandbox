package html

type ATarget string

const (
	ATargetBlank  = "_blank"
	ATargetSelf   = "_self"
	ATargetParent = "_parent"
	ATargetTop    = "_top"
)

type AElement struct {
	href     string
	target   ATarget
	children []any // Element | string
}

func A(href string, children ...any /* Element | string */) *AElement {
	return &AElement{href: href, children: children}
}

func (a *AElement) Target(target ATarget) *AElement {
	a.target = target
	return a
}

func (a *AElement) Render() error {
	return nil
}
