package html

import "io"

type Context struct {
	io.Writer
	elem   Element
	parent Element
}

func NewContext(elem Element, parent Element, w io.Writer) *Context {
	return &Context{w, elem, parent}
}

func (ctx *Context) WriteString(s string) (int, error) {
	return io.WriteString(ctx.Writer, s)
}

func (ctx *Context) Parent() Element {
	return ctx.parent
}

func (ctx *Context) RenderChild(child Element) error {
	return child.Render(&Context{ctx.Writer, child, ctx.elem})
}
