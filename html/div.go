package html

type Div struct {
	Content []Element
}

var _ Element = (*Div)(nil)

func (div *Div) Render(ctx *Context) error {
	if _, err := ctx.WriteString("<div>"); err != nil {
		return err
	}

	for _, elem := range div.Content {
		if err := elem.Render(ctx); err != nil {
			return err
		}
	}

	if _, err := ctx.WriteString("</div>"); err != nil {
		return err
	}

	return nil
}
