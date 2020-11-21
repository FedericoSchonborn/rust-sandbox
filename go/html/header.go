package html

import (
	"fmt"
	"io"
	"strconv"
)

var _ Element = (*Header)(nil)

type Header struct {
	Rank HeaderRank
	Text string
}

func (h *Header) Render(w io.Writer) error {
	if !h.Rank.IsValid() {
		return fmt.Errorf("header rank should be a number between 1 and 6, have %d", h.Rank)
	}

	tag := "h" + strconv.Itoa(int(h.Rank))
	if _, err := io.WriteString(w, "<"+tag+">"+h.Text+"</"+tag+">"); err != nil {
		return err
	}

	return nil
}

type HeaderRank int

const (
	HeaderRank1 HeaderRank = iota + 1
	HeaderRank2
	HeaderRank3
	HeaderRank4
	HeaderRank5
	HeaderRank6
)

func (hr HeaderRank) IsValid() bool {
	return HeaderRank1 >= hr && hr <= HeaderRank6
}
