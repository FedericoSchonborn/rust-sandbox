package html

import "io"

type Context struct {
	w io.Writer
}

func NewContext(w io.Writer) *Context {
	return &Context{w}
}

func (ctx *Context) Write(data []byte) (n int, err error) {
	return ctx.w.Write(data)
}

func (ctx *Context) WriteString(data string) (n int, err error) {
	return io.WriteString(ctx.w, data)
}
