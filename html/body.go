package html

type BodyElement struct {
	content []Element
}

func Body(content ...Element) *BodyElement {
	return &BodyElement{
		content: content,
	}
}

func (body *BodyElement) Render(ctx *Context) error {
	if _, err := ctx.WriteString("<body>"); err != nil {
		return err
	}

	for _, elem := range body.content {
		if err := ctx.RenderChild(elem); err != nil {
			return err
		}
	}

	if _, err := ctx.WriteString("</body>"); err != nil {
		return err
	}

	return nil
}
