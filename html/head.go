package html

type HeadElement struct{}

func Head() *HeadElement {
	return &HeadElement{}
}

func (head *HeadElement) Render(ctx *Context) error {
	return nil
}
