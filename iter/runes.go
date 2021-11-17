package iter

type runes struct {
	inner []rune
	len   int
	index int
}

func Runes(s string) Iterator[rune] {
	inner := []rune(s)

	return &runes{
		inner: inner,
		len:   len(inner),
		index: 0,
	}
}

func (b *runes) Next() (_ rune, ok bool) {
	if b.index >= b.len {
		return 0, false
	}
	defer func() { b.index++ }()

	return b.inner[b.index], true
}
