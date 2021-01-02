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
	return NewBuilder("body", TagKindBlock).Children(body.content).Render(ctx)
}
