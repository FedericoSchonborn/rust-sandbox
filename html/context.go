package html

import (
	"html"
	"io"
)

type Context struct {
	w      io.Writer
	elem   Element
	parent Element
}

func NewContext(elem Element, parent Element, w io.Writer) *Context {
	return &Context{w, elem, parent}
}

func (ctx *Context) Write(b []byte) error {
	_, err := ctx.w.Write(b)
	return err
}

func (ctx *Context) WriteString(s string) error {
	_, err := io.WriteString(ctx.w, s)
	return err
}

func (ctx *Context) WriteEscaped(s string) error {
	_, err := io.WriteString(ctx.w, html.EscapeString(s))
	return err
}

func (ctx *Context) Parent() Element {
	return ctx.parent
}

func (ctx *Context) RenderChild(child Element) error {
	return child.Render(&Context{ctx.w, child, ctx.elem})
}

func (ctx *Context) RenderChildren(children []Element) error {
	for _, child := range children {
		if err := ctx.RenderChild(child); err != nil {
			return err
		}
	}

	return nil
}
