package html

type PElement struct {
	text string
}

func P(text string) *PElement {
	return &PElement{
		text: text,
	}
}

func (p *PElement) Render(ctx *Context) error {
	if _, err := ctx.WriteString("<p>" + p.text + "</p>"); err != nil {
		return err
	}

	return nil
}
