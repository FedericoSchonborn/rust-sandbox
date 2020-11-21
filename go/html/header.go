package html

import (
	"fmt"
	"io"
	"strconv"
)

var _ BodyChild = (*HeaderElement)(nil)

type HeaderElement struct {
	rank int
	text string
}

func H(rank int, text string) *HeaderElement {
	return &HeaderElement{
		rank: rank,
		text: text,
	}
}

func (h *HeaderElement) Render(body *BodyElement, w io.Writer) error {
	if h.rank < 1 || h.rank > 6 {
		return fmt.Errorf("header rank should be a number between 1 and 6, have %d", h.rank)
	}

	tag := "h" + strconv.Itoa(h.rank)
	if _, err := io.WriteString(w, "<"+tag+">"+h.text+"</"+tag+">"); err != nil {
		return err
	}

	return nil
}
