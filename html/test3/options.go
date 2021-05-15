package test3

type optionChildren []Element

func Children(elems ...Element) optionChildren {
	return optionChildren(elems)
}

func (a optionChildren) bodyOption(body *BodyElement) error {
	body.children = append(body.children, a...)
	return nil
}

type optionText string

func Text(s string) optionText {
	return optionText(s)
}

func (t optionText) pOption(p *PElement) error {
	p.text = string(t)
	return nil
}
