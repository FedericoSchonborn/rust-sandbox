package html

var _ Element = (*DivElement)(nil)

type DivElement struct {
	content []Element
}

func Div(content ...Element) *DivElement {
	return &DivElement{
		content: content,
	}
}

func (div *DivElement) Render(ctx *Context) error {
	return NewBuilder("div", TagKindBlock).Children(div.content).Render(ctx)
}
