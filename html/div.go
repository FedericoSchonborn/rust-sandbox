package html

var _ Element = (*DivElement)(nil)

type DivElement struct {
	content []Element
	class   string
}

func Div(content ...Element) *DivElement {
	return &DivElement{
		content: content,
	}
}

func (div *DivElement) Render(ctx *Context) error {
	openTag := "<div"
	if div.class != "" {
		openTag += " class=\"" + div.class + "\""
	}
	openTag += ">"

	if _, err := ctx.WriteString(openTag); err != nil {
		return err
	}

	for _, elem := range div.content {
		if err := ctx.RenderChild(elem); err != nil {
			return err
		}
	}

	if _, err := ctx.WriteString("</div>"); err != nil {
		return err
	}

	return nil
}
