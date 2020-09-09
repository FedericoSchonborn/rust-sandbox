package html

type Header struct {
	Rank HeaderRank
	Text string
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

var _ Element = (*Header)(nil)

func (h *Header) Render(ctx *Context) error {
	return nil
}
