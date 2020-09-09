package html

import (
	"io"
)

type Context struct {
	w io.Writer
}

func NewContext(w io.Writer) *Context {
	return &Context{
		w: w,
	}
}
