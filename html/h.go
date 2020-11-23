package html

import (
	"fmt"
	"strconv"
)

var _ Element = (*HElement)(nil)

type HElement struct {
	rank int
	text string
}

func h(rank int, text string) *HElement {
	return &HElement{
		rank: rank,
		text: text,
	}
}

func H1(text string) *HElement {
	return h(1, text)
}

func H2(text string) *HElement {
	return h(2, text)
}

func H3(text string) *HElement {
	return h(3, text)
}

func H4(text string) *HElement {
	return h(4, text)
}

func H5(text string) *HElement {
	return h(5, text)
}

func H6(text string) *HElement {
	return h(6, text)
}

func (h *HElement) Render(ctx *Context) error {
	if h.rank < 1 || h.rank > 6 {
		return fmt.Errorf("header rank should be a number between 1 and 6, have %d", h.rank)
	}

	tag := "h" + strconv.Itoa(h.rank)
	if _, err := ctx.WriteString("<" + tag + ">" + h.text + "</" + tag + ">"); err != nil {
		return err
	}

	return nil
}
