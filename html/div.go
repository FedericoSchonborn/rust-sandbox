package html

type Div struct {
	Content []Element
}

var _ Element = (*Div)(nil)

func (div *Div) Render(ctx *Context) error {
	return nil
}
